package repository

import "zisui-suki-blog/domain/model"

type TagRepository interface {
	FindByTagName(model.TagName) (bool, *model.Tag, error)
	Register(*model.Tag) error
}
