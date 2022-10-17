package presenter

import (
	"zisui-suki-blog/usecase/apperr"
	"zisui-suki-blog/usecase/userapp"

	"github.com/labstack/echo/v4"
)

type UserPresenter struct {
	c echo.Context
}

func NewUserPresenter(
	c echo.Context,
) userapp.UserOutputPort {
	return &UserPresenter{
		c: c,
	}
}

func (u *UserPresenter) RespondUser(response *userapp.UserResponse) error {
	
}

func (u *UserPresenter) RespondLoginSuccess(response *userapp.UserResponse) error {

}

func (u *UserPresenter) RespondLogoutSuccess(response *userapp.UserResponse) error {
	
}

func (u *UserPresenter) RespondSignupSuccess(response *userapp.UserResponse) error {
	
}

func (u *UserPresenter) RespondSignupFailure() error {
	
}

func (u *UserPresenter) RespondAuthenticationFailure() error {
	
}

func (u *UserPresenter) RespondError(response *apperr.ErrorResponse) error {
	
}