package repository

import "cook-blog/domain/model"

/* blog repository */
// * PK で Find するメソッドは、返り値の1つ目が存在フラグ
type BlogRepository interface {
	FindById(model.BlogId) (bool, *model.Blog, error)
	FindByUserId(model.UserId) ([]*model.Blog, error)
	Register(*model.Blog) error
	Update(*model.Blog) (*model.Blog, error)
	Delete(*model.Blog) error
}
