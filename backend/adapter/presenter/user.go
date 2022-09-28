package presenter

import (
	"cook-blog/usecase/userapp"

	"github.com/gin-gonic/gin"
)

type UserPresenter struct {
	c *gin.Context
}

func NewUserPresenter(
	c *gin.Context,
) userapp.UserOutputPort {
	return &UserPresenter{
		c: c,
	}
}

func (u *UserPresenter) RespondUser(user userapp.UserResponse) {
	
}