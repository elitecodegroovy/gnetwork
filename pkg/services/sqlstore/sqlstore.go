package sqlstore

import (
	"fmt"
	"github.com/elitecodegroovy/gnetwork/pkg/bus"
	"github.com/elitecodegroovy/gnetwork/pkg/infra/localcache"
	"github.com/elitecodegroovy/gnetwork/pkg/infra/log"
	"github.com/elitecodegroovy/gnetwork/pkg/registry"
	"github.com/elitecodegroovy/gnetwork/pkg/services/annotations"
	"github.com/elitecodegroovy/gnetwork/pkg/services/sqlstore/migrator"
	"github.com/elitecodegroovy/gnetwork/pkg/services/sqlstore/sqlutil"
	"github.com/elitecodegroovy/gnetwork/pkg/setting"
	"github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

const ContextSessionName = "db-session"

var (
	x       *xorm.Engine
	dialect migrator.Dialect

	sqlog log.Logger = log.New("sqlstore")
)

func init() {
	// This change will make xorm use an empty default schema for postgres and
	// by that mimic the functionality of how it was functioning before
	// xorm's changes above.
	xorm.DefaultPostgresSchema = ""

	registry.Register(&registry.Descriptor{
		Name:         "SqlStore",
		Instance:     &SqlStore{},
		InitPriority: registry.High,
	})
	fmt.Println("Initialized sqlstore ....")
}

type SqlStore struct {
	Cfg          *setting.Cfg             `inject:""`
	Bus          bus.Bus                  `inject:""`
	CacheService *localcache.CacheService `inject:""`

	dbCfg   DatabaseConfig
	Dialect migrator.Dialect
	engine  *xorm.Engine
	log     log.Logger
}

type DatabaseConfig struct {
	Type             string
	Host             string
	Name             string
	User             string
	Pwd              string
	Path             string
	SslMode          string
	CaCertPath       string
	ClientKeyPath    string
	ClientCertPath   string
	ServerCertName   string
	ConnectionString string
	MaxOpenConn      int
	MaxIdleConn      int
	ConnMaxLifetime  int
	CacheMode        string
	UrlQueryParams   map[string][]string
}

func (ss *SqlStore) Init() error {
	ss.log = log.New("sqlstore")
	ss.readConfig()

	engine, err := ss.getEngine()

	if err != nil {
		return fmt.Errorf("Fail to connect to database: %v", err)
	}

	ss.engine = engine

	// temporarily still set global var
	x = engine
	dialect = ss.Dialect

	// Init repo instances
	annotations.SetRepository(&SqlAnnotationRepo{})
	ss.Bus.SetTransactionManager(ss)

	// Register handlers
	ss.addUserQueryAndCommandHandlers()

	return nil
}

func (ss *SqlStore) readConfig() {
	sec := ss.Cfg.Raw.Section("database")

	cfgURL := sec.Key("url").String()
	if len(cfgURL) != 0 {
		dbURL, _ := url.Parse(cfgURL)
		ss.dbCfg.Type = dbURL.Scheme
		ss.dbCfg.Host = dbURL.Host

		pathSplit := strings.Split(dbURL.Path, "/")
		if len(pathSplit) > 1 {
			ss.dbCfg.Name = pathSplit[1]
		}

		userInfo := dbURL.User
		if userInfo != nil {
			ss.dbCfg.User = userInfo.Username()
			ss.dbCfg.Pwd, _ = userInfo.Password()
		}

		ss.dbCfg.UrlQueryParams = dbURL.Query()
	} else {
		ss.dbCfg.Type = sec.Key("type").String()
		ss.dbCfg.Host = sec.Key("host").String()
		ss.dbCfg.Name = sec.Key("name").String()
		ss.dbCfg.User = sec.Key("user").String()
		ss.dbCfg.ConnectionString = sec.Key("connection_string").String()
		ss.dbCfg.Pwd = sec.Key("password").String()
	}

	ss.dbCfg.MaxOpenConn = sec.Key("max_open_conn").MustInt(0)
	ss.dbCfg.MaxIdleConn = sec.Key("max_idle_conn").MustInt(2)
	ss.dbCfg.ConnMaxLifetime = sec.Key("conn_max_lifetime").MustInt(14400)

	ss.dbCfg.SslMode = sec.Key("ssl_mode").String()
	ss.dbCfg.CaCertPath = sec.Key("ca_cert_path").String()
	ss.dbCfg.ClientKeyPath = sec.Key("client_key_path").String()
	ss.dbCfg.ClientCertPath = sec.Key("client_cert_path").String()
	ss.dbCfg.ServerCertName = sec.Key("server_cert_name").String()
	ss.dbCfg.Path = sec.Key("path").MustString("data/grafana.db")

	ss.dbCfg.CacheMode = sec.Key("cache_mode").MustString("private")
}

func (ss *SqlStore) buildConnectionString() (string, error) {
	cnnstr := ss.dbCfg.ConnectionString

	// special case used by integration tests
	if cnnstr != "" {
		return cnnstr, nil
	}

	switch ss.dbCfg.Type {
	case migrator.MYSQL:
		protocol := "tcp"
		if strings.HasPrefix(ss.dbCfg.Host, "/") {
			protocol = "unix"
		}

		cnnstr = fmt.Sprintf("%s:%s@%s(%s)/%s?collation=utf8mb4_unicode_ci&allowNativePasswords=true",
			ss.dbCfg.User, ss.dbCfg.Pwd, protocol, ss.dbCfg.Host, ss.dbCfg.Name)

		if ss.dbCfg.SslMode == "true" || ss.dbCfg.SslMode == "skip-verify" {
			tlsCert, err := makeCert(ss.dbCfg)
			if err != nil {
				return "", err
			}
			mysql.RegisterTLSConfig("custom", tlsCert)
			cnnstr += "&tls=custom"
		}

		cnnstr += ss.buildExtraConnectionString('&')

	case migrator.SQLITE:
		// special case for tests
		if !filepath.IsAbs(ss.dbCfg.Path) {
			ss.dbCfg.Path = filepath.Join(ss.Cfg.DataPath, ss.dbCfg.Path)
		}
		os.MkdirAll(path.Dir(ss.dbCfg.Path), os.ModePerm)
		cnnstr = fmt.Sprintf("file:%s?cache=%s&mode=rwc", ss.dbCfg.Path, ss.dbCfg.CacheMode)
		cnnstr += ss.buildExtraConnectionString('&')
	default:
		return "", fmt.Errorf("Unknown database type: %s", ss.dbCfg.Type)
	}

	return cnnstr, nil
}

func (ss *SqlStore) getEngine() (*xorm.Engine, error) {
	connectionString, err := ss.buildConnectionString()

	if err != nil {
		return nil, err
	}

	sqlog.Info("Connecting to DB", "dbtype", ss.dbCfg.Type)
	engine, err := xorm.NewEngine(ss.dbCfg.Type, connectionString)
	if err != nil {
		return nil, err
	}

	engine.SetMaxOpenConns(ss.dbCfg.MaxOpenConn)
	engine.SetMaxIdleConns(ss.dbCfg.MaxIdleConn)
	engine.SetConnMaxLifetime(time.Second * time.Duration(ss.dbCfg.ConnMaxLifetime))

	// configure sql logging
	debugSql := ss.Cfg.Raw.Section("database").Key("log_queries").MustBool(false)
	if !debugSql {
		engine.SetLogger(&xorm.DiscardLogger{})
	} else {
		engine.SetLogger(NewXormLogger(log.LvlInfo, log.New("sqlstore.xorm")))
		engine.ShowSQL(true)
		engine.ShowExecTime(true)
	}

	return engine, nil
}

func (ss *SqlStore) buildExtraConnectionString(sep rune) string {
	if ss.dbCfg.UrlQueryParams == nil {
		return ""
	}

	var sb strings.Builder
	for key, values := range ss.dbCfg.UrlQueryParams {
		for _, value := range values {
			sb.WriteRune(sep)
			sb.WriteString(key)
			sb.WriteRune('=')
			sb.WriteString(value)
		}
	}
	return sb.String()
}

// Interface of arguments for testing db
type ITestDB interface {
	Helper()
	Fatalf(format string, args ...interface{})
}

// InitTestDB initiliaze test DB
func InitTestDB(t ITestDB) *SqlStore {
	t.Helper()
	sqlstore := &SqlStore{}
	//sqlstore.skipEnsureAdmin = true
	sqlstore.Bus = bus.New()

	dbType := migrator.SQLITE

	// environment variable present for test db?
	if db, present := os.LookupEnv("GRAFANA_TEST_DB"); present {
		dbType = db
	}

	// set test db config
	sqlstore.Cfg = setting.NewCfg()
	sec, _ := sqlstore.Cfg.Raw.NewSection("database")
	sec.NewKey("type", dbType)

	switch dbType {
	case "mysql":
		sec.NewKey("connection_string", sqlutil.TestDB_Mysql.ConnStr)
	case "postgres":
		sec.NewKey("connection_string", sqlutil.TestDB_Postgres.ConnStr)
	default:
		sec.NewKey("connection_string", sqlutil.TestDB_Sqlite3.ConnStr)
	}

	// need to get engine to clean db before we init
	engine, err := xorm.NewEngine(dbType, sec.Key("connection_string").String())
	if err != nil {
		t.Fatalf("Failed to init test database: %v", err)
	}

	sqlstore.Dialect = migrator.NewDialect(engine)

	// temp global var until we get rid of global vars
	dialect = sqlstore.Dialect

	if err := dialect.CleanDB(); err != nil {
		t.Fatalf("Failed to clean test db %v", err)
	}

	if err := sqlstore.Init(); err != nil {
		t.Fatalf("Failed to init test database: %v", err)
	}

	sqlstore.engine.DatabaseTZ = time.UTC
	sqlstore.engine.TZLocation = time.UTC

	return sqlstore
}
