package controller

import (
	"context"
	"zisui-suki-blog/adapter/gateway"
	"zisui-suki-blog/adapter/presenter"
	"zisui-suki-blog/ent"
	"zisui-suki-blog/usecase/userapp"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	client *ent.Client
}

func NewUserController(client *ent.Client) *UserController {
	return &UserController{
		client: client,
	}
}

func (u *UserController) NameToUserId(ctx *context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		/* get request parameters */
		userName := c.QueryParam("name")
		request := &userapp.UserNameToUserIdRequest{
			Name: userName,
		}
		
		/* return response */
		return u.inputPort(c, ctx).NameToUserId(request)
	}
}

func (u *UserController) Login(ctx *context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		/* get request parameters */
		request := &userapp.UserLoginRequest{}
		err := c.Bind(request)
		if err != nil {
			return err
		}

		/* return response */
		return u.inputPort(c, ctx).Login(request)
	}
}

func (u *UserController) SignUp(ctx *context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		/* get request parameters */
		request := &userapp.UserSignUpRequest{}
		err := c.Bind(request)
		if err != nil {
			return err
		}

		/* return response */
		return u.inputPort(c, ctx).SignUp(request)
	}
}

func (u *UserController) FindByToken(ctx *context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		/* get user_id from token */
		register := c.Get("user").(*jwt.Token)
		claims := register.Claims.(jwt.MapClaims)
		userId := claims["userId"].(string)
		request := userapp.UserFindByIdRequest{
			UserId: userId,
		}

		/* return response */
		return u.inputPort(c, ctx).FindById(&request)
	}
}

func (u *UserController) FindById(ctx *context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		/* get request parameters */
		userId := c.QueryParam("userId")
		request := &userapp.UserFindByIdRequest{
			UserId: userId,
		}

		/* return response */
		return u.inputPort(c, ctx).FindById(request)
	}
}

func (u *UserController) FindByName(ctx *context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		/* get request parameters */
		name := c.QueryParam("name")
		request := &userapp.UserFindByNameRequest{
			Name: name,
		}

		/* return response */
		return u.inputPort(c, ctx).FindByName(request)
	}
}

func (u *UserController) Update(ctx *context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		/* get request parameters */
		request := &userapp.UserUpdateRequest{}
		err := c.Bind(request)
		if err != nil {
			return err
		}

		/* return response */
		return u.inputPort(c, ctx).Update(request)
	}
}

func (u *UserController) UpdatePassword(ctx *context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		/* get request parameters */
		request := &userapp.UserUpdatePasswordRequest{}
		err := c.Bind(request)
		if err != nil {
			return err
		}

		/* return response */
		return u.inputPort(c, ctx).UpdatePassword(request)
	}
}

func (u *UserController) Delete(ctx *context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		/* get request parameters */
		request := &userapp.UserDeleteRequest{}
		err := c.Bind(request)
		if err != nil {
			return err
		}

		/* return response */
		return u.inputPort(c, ctx).Delete(request)
	}
}


/* create input port */
func (u *UserController) inputPort(c echo.Context, ctx *context.Context) userapp.UserInputPort {
	outputPort := presenter.NewUserPresenter(c)
	repo := gateway.NewUserGateway(u.client, ctx)
	return userapp.NewUserInteractor(outputPort, repo)
}