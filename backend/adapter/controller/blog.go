package controller

import (
	"context"
	"zisui-suki-blog/adapter/gateway"
	"zisui-suki-blog/adapter/presenter"
	"zisui-suki-blog/ent"
	"zisui-suki-blog/usecase/blogapp"

	"github.com/labstack/echo/v4"
)

type BlogController struct {
	client *ent.Client
}

func NewBlogController(client *ent.Client) *BlogController {
	return &BlogController{
		client: client,
	}
}

func (b *BlogController) Register(ctx *context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		/* get request parameters */
		request := &blogapp.BlogRegisterRequest{}
		err := c.Bind(request)
		if err != nil {
			return err
		}

		/* return response */
		return b.newBlogInputPort(c, ctx).Register(request)
	}
}

func (b *BlogController) Update(ctx *context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		/* get request parameters */
		request := &blogapp.BlogUpdateRequest{}
		err := c.Bind(request)
		if err != nil {
			return err
		}

		/* return response */
		return b.newBlogInputPort(c, ctx).Update(request)
	}
}

func (b *BlogController) FindById(ctx *context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		/* get request parameters */
		request := &blogapp.BlogFindByIdRequest{}
		err := c.Bind(request)
		if err != nil {
			return err
		}

		/* return response */
		return b.newBlogInputPort(c, ctx).FindById(request)
	}
}

func (b *BlogController) FindByTagName(ctx *context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		/* get request parameters */
		request := &blogapp.BlogFindByTagRequest{}
		err := c.Bind(request)
		if err != nil {
			return err
		}

		/* return response */
		return b.newBlogInputPort(c, ctx).FindByTagName(request)
	}
}

func (b *BlogController) FindByUserId(ctx *context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		/* get request parameters */
		request := &blogapp.BlogFindByUserIdRequest{}
		err := c.Bind(request)
		if err != nil {
			return err
		}

		/* return response */
		return b.newBlogInputPort(c, ctx).FindByUserId(request)
	}
}

func (b *BlogController) FindByUserIdAndTagName(ctx *context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		/* get request parameters */
		request := &blogapp.BlogFindByUserIdAndTagRequest{}
		err := c.Bind(request)
		if err != nil {
			return err
		}

		/* return response */
		return b.newBlogInputPort(c, ctx).FindByUserIdAndTagName(request)
	}
}

func (b *BlogController) Delete(ctx *context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		/* get request parameters */
		request := &blogapp.BlogDeleteRequest{}
		err := c.Bind(request)
		if err != nil {
			return err
		}

		/* return response */
		return b.newBlogInputPort(c, ctx).Delete(request)
	}
}

/* create input port */
func (b *BlogController) newBlogInputPort(c echo.Context, ctx *context.Context) blogapp.BlogInputPort {
	outputPort := presenter.NewBlogPresenter(c)
	repository := gateway.NewBlogGateway(b.client, ctx)
	return blogapp.NewBlogInteractor(outputPort, repository)
}
