package controller

import (
	"cook-blog/domain/repository"
	"cook-blog/usecase/userapp"

	"github.com/gin-gonic/gin"
)

type OutputFactory func(c *gin.Context) userapp.UserOutputPort
type InputFactory func(userapp.UserOutputPort, repository.UserRepository) userapp.UserInputPort
type RepositoryFactory func()

type UserController struct {
	outputFactory     OutputFactory
	inputFactory      InputFactory
	repositoryFactory RepositoryFactory
	
}

func NewUserController()

func (u *UserController) Login(c *gin.Context) {
	/*
		username = c.username
		password = c.password

		request = CreateRequest(username, password)

		userapp.Login(request)
	*/
}

func (u *UserController) Update(c *gin.Context) {
	/*

	 */
}
