package server

import (
	"errors"
	"github.com/elitecodegroovy/gnetwork/pkg/bus"
	"github.com/elitecodegroovy/gnetwork/pkg/infra/metrics"
	"github.com/elitecodegroovy/gnetwork/pkg/middleware"
	m "github.com/elitecodegroovy/gnetwork/pkg/models"
	"github.com/elitecodegroovy/gnetwork/pkg/server/dtos"
	"github.com/elitecodegroovy/gnetwork/pkg/setting"
	"net/url"
)

var (
	ErrEmailNotAllowed       = errors.New("Required email domain not fulfilled")
	ErrInvalidCredentials    = errors.New("Invalid Username or Password")
	ErrNoEmail               = errors.New("Login provider didn't return an email address")
	ErrProviderDeniedRequest = errors.New("Login provider denied login request")
	ErrSignUpNotAllowed      = errors.New("Signup is not allowed for this adapter")
	ErrTooManyLoginAttempts  = errors.New("Too many consecutive incorrect login attempts for user. Login for user temporarily blocked")
	ErrPasswordEmpty         = errors.New("No password provided")
	ErrUserDisabled          = errors.New("User is disabled")
)

func (hs *HTTPServer) LoginPost(c *m.ReqContext, cmd dtos.LoginCommand) Response {
	if setting.DisableLoginForm {
		return Error(401, "Login is disabled", nil)
	}

	authQuery := &m.LoginUserQuery{
		ReqContext: c,
		Username:   cmd.User,
		Password:   cmd.Password,
		IpAddress:  c.Req.RemoteAddr,
	}

	if err := bus.Dispatch(authQuery); err != nil {
		if err == ErrInvalidCredentials || err == ErrTooManyLoginAttempts {
			return Error(401, "Invalid username or password", err)
		}

		if err == ErrUserDisabled {
			return Error(401, "User is disabled", err)
		}

		return Error(500, "Error while trying to authenticate user", err)
	}

	user := authQuery.User

	hs.loginUserWithUser(user, c)

	result := map[string]interface{}{
		"message": "Logged in",
	}

	if redirectTo, _ := url.QueryUnescape(c.GetCookie("redirect_to")); len(redirectTo) > 0 {
		result["redirectUrl"] = redirectTo
		c.SetCookie("redirect_to", "", -1, setting.AppSubUrl+"/")
	}

	metrics.M_Api_Login_Post.Inc()

	return JSON(200, result)
}

func (hs *HTTPServer) loginUserWithUser(user *m.User, c *m.ReqContext) {
	if user == nil {
		hs.log.Error("user login with nil user")
	}

	userToken, err := hs.AuthTokenService.CreateToken(c.Req.Context(), user.Id, c.RemoteAddr(), c.Req.UserAgent())
	if err != nil {
		hs.log.Error("failed to create auth token", "error", err)
	}

	middleware.WriteSessionCookie(c, userToken.UnhashedToken, hs.Cfg.LoginMaxLifetimeDays)
}
