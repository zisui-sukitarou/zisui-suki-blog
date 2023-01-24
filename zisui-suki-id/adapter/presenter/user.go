package presenter

import (
	"net/http"
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

func (u *UserPresenter) NameToUserId(response *userapp.UserNameToUserIdResponse) error {
	return u.c.JSON(http.StatusOK, response)
}

func (u *UserPresenter) Login(response *userapp.LoginResponse) error {
	return u.c.JSON(http.StatusOK, response)
}

func (u *UserPresenter) SignUp(response *userapp.SignUpResponse) error {
	return u.c.JSON(http.StatusCreated, response)
}

func (u *UserPresenter) FindById(response *userapp.UserResponse) error {
	return u.c.JSON(http.StatusOK, response)
}

func (u *UserPresenter) FindByName(response *userapp.UserResponse) error {
	return u.c.JSON(http.StatusOK, response)
}

func (u *UserPresenter) Update(response *userapp.UpdateResponse) error {
	return u.c.JSON(http.StatusCreated, response)
}

func (u *UserPresenter) UpdatePassword(response *userapp.UpdatePasswordResponse) error {
	return u.c.JSON(http.StatusCreated, response)
}

func (u *UserPresenter) Delete() error {
	return u.c.JSON(http.StatusCreated, nil)
}

func (u *UserPresenter) RespondError(response *apperr.ErrorResponse) error {
	return u.c.JSON(http.StatusBadRequest, response)
}