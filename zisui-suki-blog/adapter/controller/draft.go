package controller

import (
	"context"
	"strconv"
	"zisui-suki-blog/adapter/gateway"
	"zisui-suki-blog/adapter/presenter"
	"zisui-suki-blog/ent"
	"zisui-suki-blog/usecase/draftapp"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type DraftController struct {
	dbClient      *ent.Client
	authApiDomain string
}

func NewDraftController(
	dbClient *ent.Client,
	authApiDomain string,
) *DraftController {
	return &DraftController{
		dbClient:      dbClient,
		authApiDomain: authApiDomain,
	}
}

func (d *DraftController) New(ctx *context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		/* get request parameters */
		request := &draftapp.DraftNewRequest{}
		err := c.Bind(request)
		if err != nil {
			return err
		}

		/* return response */
		return d.inputPort(c, ctx).New(request)
	}
}

// TODO: 書き込み系は、db内の`user_id`と一致するか確認する

func (d *DraftController) Register(ctx *context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		/* get request parameters */
		request := &draftapp.DraftRegisterRequest{}
		err := c.Bind(request)
		if err != nil {
			return err
		}

		/* get user_id from token */
		token := c.Get("user").(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)
		userId := claims["userId"].(string)
		request.UserId = userId

		/* return response */
		return d.inputPort(c, ctx).Register(request)
	}
}

func (d *DraftController) Update(ctx *context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		/* get request parameters */
		request := &draftapp.DraftUpdateRequest{}
		err := c.Bind(request)
		if err != nil {
			return err
		}

		/* get user_id from token */
		token := c.Get("user").(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)
		userId := claims["userId"].(string)
		request.UserId = userId

		/* return response */
		return d.inputPort(c, ctx).Update(request)
	}
}

func (d *DraftController) FindById(ctx *context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		/* get request parameters */
		draftId := c.QueryParam("draftId")

		request := &draftapp.DraftFindByIdRequest{
			DraftId: draftId,
		}

		/* return response */
		return d.inputPort(c, ctx).FindById(request)
	}
}

func (d *DraftController) FindByUserId(ctx *context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		/* get user_id from token */
		token := c.Get("user").(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)
		userId := claims["userId"].(string)

		/* get request parameters */
		begin, err := strconv.Atoi(c.QueryParam("begin"))
		if err != nil {
			return err
		}
		
		end, err := strconv.Atoi(c.QueryParam("end"))
		if err != nil {
			return err
		}

		request := &draftapp.DraftFindByUserIdRequest{
			UserId: userId,
			Begin:  uint(begin),
			End:    uint(end),
		}

		/* return response */
		return d.inputPort(c, ctx).FindByUserId(request)
	}
}


func (d *DraftController) Delete(ctx *context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		/* get request parameters */
		request := &draftapp.DraftDeleteRequest{}
		err := c.Bind(request)
		if err != nil {
			return err
		}

		/* return response */
		return d.inputPort(c, ctx).Delete(request)
	}
}

/* create input port */
func (d *DraftController) inputPort(c echo.Context, ctx *context.Context) draftapp.DraftInputPort {
	outputPort := presenter.NewDraftPresenter(c)
	draftRepo := gateway.NewDraftGateway(d.dbClient, ctx)
	userRepo := gateway.NewUserGateway(d.authApiDomain, c)
	tagRepo := gateway.NewTagGateway(d.dbClient, ctx)

	return draftapp.NewDraftInteractor(outputPort, draftRepo, userRepo, tagRepo)
}
