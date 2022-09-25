package repository

import "cook-blog/domain/model"

type BlogRepository interface {
	FindById(model.BlogId) *model.Blog
	FindByUserId(model.UserId) []*model.Blog
	Register(*model.Blog)
	Update(*model.Blog) *model.Blog
	Delete(*model.Blog)
}
