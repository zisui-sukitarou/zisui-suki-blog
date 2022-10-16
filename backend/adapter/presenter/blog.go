package presenter

import (
	"zisui-suki-blog/usecase/blogapp"

	"github.com/labstack/echo/v4"
)

type BlogPresenter struct {
	c echo.Context
}

func NewBlogPresenter(
	c *echo.Context,
) blogapp.BlogOutputPort {
	return &BlogPresenter{
		ctx: ctx,
	}
}
