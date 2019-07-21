package setting

import (
	"github.com/elitecodegroovy/gnetwork/pkg/infra/log"
	"gopkg.in/ini.v1"
)

type Scheme string

const (
	MYSQL  = "mysql"
	SQLITE = "sqlite3"
)

// TODO move all global vars to this struct
type Cfg struct {
	Raw *ini.File

	Logger   log.Logger
	DataPath string
	LogsPath string

	// SMTP email settings
	Smtp SmtpSettings
}

func NewCfg() *Cfg {
	return &Cfg{
		Logger: log.New("settings"),
		Raw:    ini.Empty(),
	}
}
