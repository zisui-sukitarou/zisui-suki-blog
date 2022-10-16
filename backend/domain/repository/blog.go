package repository

import "zisui-suki-blog/domain/model"

/* blog repository */
// * PK で Find するメソッドは、返り値の1つ目が存在フラグ
type BlogRepository interface {
	FindById(model.BlogId) (bool, *model.Blog, error)
	FindByUserId(model.UserId) ([]*model.Blog, error)
	FindByTagName(model.TagName) ([]*model.Blog, error)
	Register(*model.Blog) error
	RegisterTags(model.BlogId, []model.TagName) error
	Update(*model.Blog) (*model.Blog, error)
	Delete(*model.Blog) error
}
