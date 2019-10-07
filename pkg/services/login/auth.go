package login

import (
	"errors"
	"github.com/elitecodegroovy/gnetwork/pkg/bus"
	"github.com/elitecodegroovy/gnetwork/pkg/models"
	"github.com/elitecodegroovy/gnetwork/pkg/registry"
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

func init() {
	registry.RegisterService(&UserLoginService{})
}

type UserLoginService struct{}

func (l *UserLoginService) Init() error {
	bus.AddHandler("auth", AuthenticateUser)
	return nil
}

// AuthenticateUser authenticates the user via username & password
func AuthenticateUser(query *models.LoginUserQuery) error {
	if err := validateLoginAttempts(query.Username); err != nil {
		return err
	}

	if err := validatePasswordSet(query.Password); err != nil {
		return err
	}

	err := loginUsingGrafanaDB(query)
	if err == nil || (err != models.ErrUserNotFound && err != ErrInvalidCredentials && err != ErrUserDisabled) {
		return err
	}

	if err == models.ErrUserNotFound {
		return ErrInvalidCredentials
	}

	return err
}

func validatePasswordSet(password string) error {
	if len(password) == 0 || len(password) > 256 {
		return ErrPasswordEmpty
	}

	return nil
}
