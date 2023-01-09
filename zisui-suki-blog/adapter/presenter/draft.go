package presenter

import (
	"net/http"
	"zisui-suki-blog/usecase/apperr"
	"zisui-suki-blog/usecase/draftapp"

	"github.com/labstack/echo/v4"
)

type DraftPresenter struct {
	c echo.Context
}

func NewDraftPresenter(
	c echo.Context,
) draftapp.DraftOutputPort {
	return &DraftPresenter{
		c: c,
	}
}

func (d *DraftPresenter) FindById(response *draftapp.DraftResponse) error {
	return d.c.JSON(http.StatusOK, response)
}

func (d *DraftPresenter) FindByTagName(response []*draftapp.DraftOverviewResponse) error {
	return d.c.JSON(http.StatusOK, response)
}

func (d *DraftPresenter) FindByUserId(response []*draftapp.DraftOverviewResponse) error {
	return d.c.JSON(http.StatusOK, response)
}

func (d *DraftPresenter) FindByUserName(response []*draftapp.DraftOverviewResponse) error {
	return d.c.JSON(http.StatusOK, response)
}

func (d *DraftPresenter) FindByUserIdAndTagName(response []*draftapp.DraftOverviewResponse) error {
	return d.c.JSON(http.StatusOK, response)
}

func (d *DraftPresenter) FindByUserNameAndTagName(response []*draftapp.DraftOverviewResponse) error {
	return d.c.JSON(http.StatusOK, response)
}

func (d *DraftPresenter) New(response *draftapp.DraftNewResponse) error {
	return d.c.JSON(http.StatusCreated, response)
}

func (d *DraftPresenter) Register() error {
	return d.c.JSON(http.StatusCreated, nil)
}

func (d *DraftPresenter) Update() error {
	return d.c.JSON(http.StatusCreated, nil)
}

func (d *DraftPresenter) Delete() error {
	return d.c.JSON(http.StatusCreated, nil)
}

func (d *DraftPresenter) RespondErorr(response *apperr.ErrorResponse) error {
	return d.c.JSON(http.StatusBadRequest, response)
}