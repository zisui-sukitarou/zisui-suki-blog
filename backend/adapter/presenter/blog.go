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

func (b *BlogPresenter) RespondBlog(response *blogapp.BlogResponse) error {
	return b.c.JSON(http.StatusOK, response)
}

func (b *BlogPresenter) RespondBlogOverviews(response []*blogapp.BlogOverviewResponse) error {
	return b.c.JSON(http.StatusOK, response)
}

func (b *BlogPresenter) RespondErorr(response *apperr.ErrorResponse) error {
	return b.c.JSON(http.StatusBadRequest, response)
}