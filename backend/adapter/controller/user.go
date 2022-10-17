package controller

import (
	"context"
	"zisui-suki-blog/adapter/gateway"
	"zisui-suki-blog/adapter/presenter"
	"zisui-suki-blog/ent"
	"zisui-suki-blog/usecase/userapp"

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


/* create input port */
func (u *UserController) inputPort(c echo.Context, ctx *context.Context) userapp.UserInputPort {
	outputPort := presenter.NewUserPresenter(c)
	repo := gateway.NewUserGateway(u.client, ctx)
	return userapp.NewUserInteractor(outputPort, repo)
}