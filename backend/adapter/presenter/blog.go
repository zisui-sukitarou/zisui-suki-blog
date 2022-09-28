package presenter

import (
	"cook-blog/usecase/blogapp"

	"github.com/gin-gonic/gin"
)

type BlogPresenter struct {
	c *gin.Context
}

func NewBlogPresenter(
	c *gin.Context,
) blogapp.BlogOutputPort {
	return &BlogPresenter{
		c: c,
	}
}

