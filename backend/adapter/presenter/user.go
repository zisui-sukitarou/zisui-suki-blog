package presenter

import (
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

func (u *UserPresenter) RespondUser(user userapp.UserResponse) {
	
}

func (u *UserPresenter) RespondLoginSuccess(user userapp.UserResponse) {

}

func (u *UserPresenter) RespondLogoutSuccess(user userapp.UserResponse) {
	
}

func (u *UserPresenter) RespondSignupSuccess(user userapp.UserResponse) {
	
}

func (u *UserPresenter) RespondSignupFailure() {
	
}

func (u *UserPresenter) RespondAuthenticationFailure() {
	
}

func (u *UserPresenter) RespondError(err error) {
	
}