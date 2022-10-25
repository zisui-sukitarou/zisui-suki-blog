package presenter

import (
	"net/http"
	"zisui-suki-blog/usecase/blogapp"
	"zisui-suki-blog/usecase/apperr"

	"github.com/labstack/echo/v4"
)

type BlogPresenter struct {
	c echo.Context
}

func NewBlogPresenter(
	c echo.Context,
) blogapp.BlogOutputPort {
	return &BlogPresenter{
		c: c,
	}
}

func (b *BlogPresenter) FindById(response *blogapp.BlogResponse) error {
	return b.c.JSON(http.StatusOK, response)
}

func (b *BlogPresenter) FindByTagName(response []*blogapp.BlogOverviewResponse) error {
	return b.c.JSON(http.StatusOK, response)
}

func (b *BlogPresenter) FindByUserId(response []*blogapp.BlogOverviewResponse) error {
	return b.c.JSON(http.StatusOK, response)
}

func (b *BlogPresenter) FindByUserIdAndTagName(response []*blogapp.BlogOverviewResponse) error {
	return b.c.JSON(http.StatusOK, response)
}

func (b *BlogPresenter) Register() error {
	return b.c.JSON(http.StatusCreated, nil)
}

func (b *BlogPresenter) Update(response *blogapp.BlogResponse) error {
	return b.c.JSON(http.StatusCreated, response)
}

func (b *BlogPresenter) Delete() error {
	return b.c.JSON(http.StatusCreated, nil)
}

func (b *BlogPresenter) RespondErorr(response *apperr.ErrorResponse) error {
	return b.c.JSON(http.StatusBadRequest, response)
}