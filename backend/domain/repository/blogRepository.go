package repository

import "cook-blog/domain/model/blog"

type BlogRepository interface {
	Save(blog model.Blog) ()
	FindById(id model.BlogId) (*model.Blog)
	Update(blog model.Blog) (*model.Blog)
	Delete(blog model.Blog) ()
}