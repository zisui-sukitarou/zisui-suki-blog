package repository

import "cook-blog/domain/model"

type UserRepository interface {
	FindById(model.UserId) *model.User
	FindByUserName(model.UserName) *model.User
	FindByEmail(model.UserEmail) *model.User
	Register(*model.User)
	Update(*model.User)
	Delete(*model.User)
}