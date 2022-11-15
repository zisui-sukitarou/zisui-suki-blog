package controller

import (
	"context"
	"strconv"
	"zisui-suki-blog/adapter/gateway"
	"zisui-suki-blog/adapter/presenter"
	"zisui-suki-blog/ent"
	"zisui-suki-blog/usecase/blogapp"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type BlogController struct {
	dbClient      *ent.Client
	authApiDomain string
}

func NewBlogController(
	dbClient *ent.Client,
	authApiDomain string,
) *BlogController {
	return &BlogController{
		dbClient:      dbClient,
		authApiDomain: authApiDomain,
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

		/* get user_id from token */
		register := c.Get("user").(*jwt.Token)
		claims := register.Claims.(jwt.MapClaims)
		userId := claims["userId"].(string)
		request.UserId = userId

		/* return response */
		return b.inputPort(c, ctx).Register(request)
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
		return b.inputPort(c, ctx).Update(request)
	}
}

func (b *BlogController) FindById(ctx *context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		/* get request parameters */
		blogId := c.QueryParam("blogId")

		request := &blogapp.BlogFindByIdRequest{
			BlogId: blogId,
		}

		/* return response */
		return b.inputPort(c, ctx).FindById(request)
	}
}

func (b *BlogController) FindByTagName(ctx *context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		/* get request parameters */
		tagName := c.QueryParam("tagName")
		begin, err := strconv.Atoi(c.QueryParam("begin"))
		if err != nil {
			return err
		}
		end, err := strconv.Atoi(c.QueryParam("end"))
		if err != nil {
			return err
		}

		request := &blogapp.BlogFindByTagRequest{
			TagName: tagName,
			Begin:   uint(begin),
			End:     uint(end),
		}

		/* return response */
		return b.inputPort(c, ctx).FindByTagName(request)
	}
}

func (b *BlogController) FindByUserId(ctx *context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		/* get request parameters */
		userId := c.QueryParam("userId")
		begin, err := strconv.Atoi(c.QueryParam("begin"))
		if err != nil {
			return err
		}
		end, err := strconv.Atoi(c.QueryParam("end"))
		if err != nil {
			return err
		}

		request := &blogapp.BlogFindByUserIdRequest{
			UserId: userId,
			Begin:  uint(begin),
			End:    uint(end),
		}

		/* return response */
		return b.inputPort(c, ctx).FindByUserId(request)
	}
}

func (b *BlogController) FindByUserName(ctx *context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		/* get request parameters */
		userName := c.QueryParam("userName")
		begin, err := strconv.Atoi(c.QueryParam("begin"))
		if err != nil {
			return err
		}
		end, err := strconv.Atoi(c.QueryParam("end"))
		if err != nil {
			return err
		}

		request := &blogapp.BlogFindByUserNameRequest{
			UserName: userName,
			Begin:    uint(begin),
			End:      uint(end),
		}

		/* return response */
		return b.inputPort(c, ctx).FindByUserName(request)
	}
}

func (b *BlogController) FindByUserIdAndTagName(ctx *context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		/* get request parameters */
		tagName := c.QueryParam("tagName")
		userId := c.QueryParam("userId")
		begin, err := strconv.Atoi(c.QueryParam("begin"))
		if err != nil {
			return err
		}
		end, err := strconv.Atoi(c.QueryParam("end"))
		if err != nil {
			return err
		}

		request := &blogapp.BlogFindByUserIdAndTagRequest{
			TagName: tagName,
			UserId:  userId,
			Begin:   uint(begin),
			End:     uint(end),
		}

		/* return response */
		return b.inputPort(c, ctx).FindByUserIdAndTagName(request)
	}
}

func (b *BlogController) FindByUserNameAndTagName(ctx *context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		/* get request parameters */
		userName := c.QueryParam("userName")
		tagName := c.QueryParam("tagName")
		begin, err := strconv.Atoi(c.QueryParam("begin"))
		if err != nil {
			return err
		}
		end, err := strconv.Atoi(c.QueryParam("end"))
		if err != nil {
			return err
		}

		request := &blogapp.BlogFindByUserNameAndTagRequest{
			UserName: userName,
			TagName:  tagName,
			Begin:    uint(begin),
			End:      uint(end),
		}

		/* return response */
		return b.inputPort(c, ctx).FindByUserNameAndTagName(request)
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
		return b.inputPort(c, ctx).Delete(request)
	}
}

/* create input port */
func (b *BlogController) inputPort(c echo.Context, ctx *context.Context) blogapp.BlogInputPort {
	outputPort := presenter.NewBlogPresenter(c)
	blogRepo := gateway.NewBlogGateway(b.dbClient, ctx)
	userRepo := gateway.NewUserGateway(b.authApiDomain, c)
	tagRepo := gateway.NewTagGateway(b.dbClient, ctx)

	return blogapp.NewBlogInteractor(outputPort, blogRepo, userRepo, tagRepo)
}
