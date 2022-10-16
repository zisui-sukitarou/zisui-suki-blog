package repository

import "zisui-suki-blog/domain/model"

/**** blog repository ****
	* PK で Find するメソッドは、返り値の1つ目が存在フラグ
	* Blog 配列が返ってくるメソッドの Blog.Content は空文字列（BlogOverviewResponse に不要なため）
*/ 
type BlogRepository interface {
	FindById(model.BlogId) (bool, *model.Blog, error)
	FindByUserId(userId model.UserId, begin uint, end uint) ([]*model.Blog, error) // Blog.Content は空文字列
	FindByTagName(tagName model.TagName, begin uint, end uint) ([]*model.Blog, error) // Blog.Content は空文字列
	FindByUserIdAndTagName(userId model.UserId, tagName model.TagName, begin uint, end uint) ([]*model.Blog, error) // Blog.Content は空文字列
	Register(*model.Blog) error
	RegisterTags(model.BlogId, []model.TagName) error
	Update(*model.Blog) (*model.Blog, error)
	Delete(*model.Blog) error
}
