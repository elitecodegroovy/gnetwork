package server

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/elitecodegroovy/gnetwork/pkg/bus"
	"github.com/elitecodegroovy/gnetwork/pkg/components/simplejson"
	"github.com/elitecodegroovy/gnetwork/pkg/infra/localcache"
	"github.com/elitecodegroovy/gnetwork/pkg/infra/log"
	"github.com/elitecodegroovy/gnetwork/pkg/middleware"
	"github.com/elitecodegroovy/gnetwork/pkg/models"
	"github.com/elitecodegroovy/gnetwork/pkg/registry"
	"github.com/elitecodegroovy/gnetwork/pkg/routing"
	httpstatic "github.com/elitecodegroovy/gnetwork/pkg/server/static"
	"github.com/elitecodegroovy/gnetwork/pkg/services/quota"
	"github.com/elitecodegroovy/gnetwork/pkg/services/remotecache"
	"github.com/elitecodegroovy/gnetwork/pkg/setting"
	"gopkg.in/macaron.v1"
	"net"
	"net/http"
	"os"
	"path"
	"time"
)

func init() {
	registry.Register(&registry.Descriptor{
		Name:         "HTTPServer",
		Instance:     &HTTPServer{},
		InitPriority: registry.High,
	})
}

type HTTPServer struct {
	log     log.Logger
	macaron *macaron.Macaron
	context context.Context
	//streamManager *live.StreamManager
	httpSrv *http.Server

	RouteRegister routing.RouteRegister `inject:""`
	Bus           bus.Bus               `inject:""`

	Cfg *setting.Cfg `inject:""`

	CacheService *localcache.CacheService `inject:""`

	AuthTokenService   models.UserTokenService  `inject:""`
	QuotaService       *quota.QuotaService      `inject:""`
	RemoteCacheService *remotecache.RemoteCache `inject:""`
	//ProvisioningService ProvisioningService      `inject:""`
	//Login               *login.LoginService      `inject:""`
}

func (hs *HTTPServer) Init() error {
	hs.log = log.New("http.server")

	hs.macaron = hs.newMacaron()
	hs.registerRoutes()

	return nil
}

func (hs *HTTPServer) newMacaron() *macaron.Macaron {
	macaron.Env = setting.Env
	m := macaron.New()

	// automatically set HEAD for every GET
	m.SetAutoHead(true)

	return m
}

func (hs *HTTPServer) applyRoutes() {
	// start with middlewares & static routes
	hs.addMiddlewaresAndStaticRoutes()
	// then add view routes & api routes
	hs.RouteRegister.Register(hs.macaron)
	// then custom app proxy routes
	//hs.initAppPluginRoutes(hs.macaron)
	// lastly not found route
	hs.macaron.NotFound(hs.NotFoundHandler)
}

func (hs *HTTPServer) addMiddlewaresAndStaticRoutes() {
	m := hs.macaron

	m.Use(middleware.Logger())

	if setting.EnableGzip {
		m.Use(middleware.Gziper())
	}

	m.Use(middleware.Recovery())

	//plugin
	//for _, route := range plugins.StaticRoutes {
	//	pluginRoute := path.Join("/public/plugins/", route.PluginId)
	//	hs.log.Debug("Plugins: Adding route", "route", pluginRoute, "dir", route.Directory)
	//	hs.mapStatic(hs.macaron, route.Directory, "", pluginRoute)
	//}

	hs.mapStatic(m, setting.StaticRootPath, "build", "public/build")
	hs.mapStatic(m, setting.StaticRootPath, "", "public")
	hs.mapStatic(m, setting.StaticRootPath, "robots.txt", "robots.txt")

	if setting.ImageUploadProvider == "local" {
		hs.mapStatic(m, hs.Cfg.ImagesDir, "", "/public/img/attachments")
	}

	m.Use(middleware.AddDefaultResponseHeaders())

	if setting.ServeFromSubPath && setting.AppSubUrl != "" {
		m.SetURLPrefix(setting.AppSubUrl)
	}

	m.Use(macaron.Renderer(macaron.RenderOptions{
		Directory:  path.Join(setting.StaticRootPath, "views"),
		IndentJSON: macaron.Env != macaron.PROD,
		Delims:     macaron.Delims{Left: "[[", Right: "]]"},
	}))

	//health for DB connection
	m.Use(hs.healthHandler)

	m.Use(middleware.GetContextHandler(
		hs.AuthTokenService,
		hs.RemoteCacheService,
	))
	//m.Use(middleware.OrgRedirect())

	m.Use(middleware.HandleNoCacheHeader())
}

func (hs *HTTPServer) healthHandler(ctx *macaron.Context) {
	notHeadOrGet := ctx.Req.Method != http.MethodGet && ctx.Req.Method != http.MethodHead
	if notHeadOrGet || ctx.Req.URL.Path != "/api/health" {
		return
	}

	data := simplejson.New()
	data.Set("database", "ok")
	data.Set("version", setting.BuildVersion)
	data.Set("commit", setting.BuildCommit)

	if err := bus.Dispatch(&models.GetDBHealthQuery{}); err != nil {
		data.Set("database", "failing")
		ctx.Resp.Header().Set("Content-Type", "application/json; charset=UTF-8")
		ctx.Resp.WriteHeader(503)
	} else {
		ctx.Resp.Header().Set("Content-Type", "application/json; charset=UTF-8")
		ctx.Resp.WriteHeader(200)
	}

	dataBytes, _ := data.EncodePretty()
	ctx.Resp.Write(dataBytes)
}

func (hs *HTTPServer) mapStatic(m *macaron.Macaron, rootDir string, dir string, prefix string) {
	headers := func(c *macaron.Context) {
		c.Resp.Header().Set("Cache-Control", "public, max-age=3600")
	}

	if prefix == "public/build" {
		headers = func(c *macaron.Context) {
			c.Resp.Header().Set("Cache-Control", "public, max-age=31536000")
		}
	}

	if setting.Env == setting.DEV {
		headers = func(c *macaron.Context) {
			c.Resp.Header().Set("Cache-Control", "max-age=0, must-revalidate, no-cache")
		}
	}

	m.Use(httpstatic.Static(
		path.Join(rootDir, dir),
		httpstatic.StaticOptions{
			SkipLogging: true,
			Prefix:      prefix,
			AddHeaders:  headers,
		},
	))
}

func (hs *HTTPServer) metricsEndpointBasicAuthEnabled() bool {
	return hs.Cfg.MetricsEndpointBasicAuthUsername != "" && hs.Cfg.MetricsEndpointBasicAuthPassword != ""
}

func (hs *HTTPServer) listenAndServeTLS(certfile, keyfile string) error {
	if certfile == "" {
		return fmt.Errorf("cert_file cannot be empty when using HTTPS")
	}

	if keyfile == "" {
		return fmt.Errorf("cert_key cannot be empty when using HTTPS")
	}

	if _, err := os.Stat(setting.CertFile); os.IsNotExist(err) {
		return fmt.Errorf(`Cannot find SSL cert_file at %v`, setting.CertFile)
	}

	if _, err := os.Stat(setting.KeyFile); os.IsNotExist(err) {
		return fmt.Errorf(`Cannot find SSL key_file at %v`, setting.KeyFile)
	}

	tlsCfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_RSA_WITH_AES_128_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
			tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		},
	}

	hs.httpSrv.TLSConfig = tlsCfg
	hs.httpSrv.TLSNextProto = make(map[string]func(*http.Server, *tls.Conn, http.Handler))

	return hs.httpSrv.ListenAndServeTLS(setting.CertFile, setting.KeyFile)
}

func (hs *HTTPServer) Run(ctx context.Context) error {
	var err error

	hs.context = ctx

	//url mapping handler
	hs.applyRoutes()

	listenAddr := fmt.Sprintf("%s:%s", setting.HttpAddr, setting.HttpPort)
	hs.log.Info("HTTP Server Listen", "address", listenAddr, "protocol", setting.Protocol, "subUrl", setting.AppSubUrl, "socket", setting.SocketPath)

	hs.httpSrv = &http.Server{Addr: listenAddr, Handler: hs.macaron}

	// handle http shutdown on server context done
	go func() {
		<-ctx.Done()
		// Hacky fix for race condition between ListenAndServe and Shutdown
		time.Sleep(time.Millisecond * 100)
		if err := hs.httpSrv.Shutdown(context.Background()); err != nil {
			hs.log.Error("Failed to shutdown server", "error", err)
		}
	}()

	switch setting.Protocol {
	case setting.HTTP:
		err = hs.httpSrv.ListenAndServe()
		if err == http.ErrServerClosed {
			hs.log.Debug("server was shutdown gracefully")
			return nil
		}
	case setting.HTTPS:
		err = hs.listenAndServeTLS(setting.CertFile, setting.KeyFile)
		if err == http.ErrServerClosed {
			hs.log.Debug("server was shutdown gracefully")
			return nil
		}
	case setting.SOCKET:
		ln, err := net.ListenUnix("unix", &net.UnixAddr{Name: setting.SocketPath, Net: "unix"})
		if err != nil {
			hs.log.Debug("server was shutdown gracefully")
			return nil
		}

		// Make socket writable by group
		os.Chmod(setting.SocketPath, 0660)

		err = hs.httpSrv.Serve(ln)
		if err != nil {
			hs.log.Debug("server was shutdown gracefully")
			return nil
		}
	default:
		hs.log.Error("Invalid protocol", "protocol", setting.Protocol)
		err = errors.New("Invalid Protocol")
	}

	hs.log.Info("server starts with ", setting.Protocol, " protocol")
	return err
}
