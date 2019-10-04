package server

import (
	"github.com/elitecodegroovy/gnetwork/pkg/middleware"
	"github.com/elitecodegroovy/gnetwork/pkg/routing"
	"github.com/elitecodegroovy/gnetwork/pkg/server/dtos"
	"github.com/go-macaron/binding"
)

func (hs *HTTPServer) registerRoutes() {
	//reqSignedIn := middleware.ReqSignedIn
	reqGrafanaAdmin := middleware.ReqGrafanaAdmin
	//reqEditorRole := middleware.ReqEditorRole
	//reqOrgAdmin := middleware.ReqOrgAdmin
	//reqCanAccessTeams := middleware.AdminOrFeatureEnabled(hs.Cfg.EditorsCanAdmin)
	//redirectFromLegacyDashboardURL := middleware.RedirectFromLegacyDashboardURL()
	//redirectFromLegacyDashboardSoloURL := middleware.RedirectFromLegacyDashboardSoloURL()
	quota := middleware.Quota(hs.QuotaService)
	bind := binding.Bind

	r := hs.RouteRegister

	// not logged in views
	r.Get("/", func() string {
		return "Macaron Web Framework!"
	})
	// admin api
	r.Group("/api/admin", func(adminRoute routing.RouteRegister) {
		adminRoute.Get("/settings", AdminGetSettings)
		adminRoute.Post("/users", bind(dtos.AdminCreateUserForm{}), AdminCreateUser)
	}, reqGrafanaAdmin)

	r.Post("/login", quota("session"), bind(dtos.LoginCommand{}), Wrap(hs.LoginPost))

	r.Get("/urlReq", urlHandler)

	r.Post("/upload", uploadFile)

	r.Get("/setting", AdminGetSettings)
	//r.Get("/api", reqSignedIn)

}
