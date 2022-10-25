package repository

import (
	"zisui-suki-blog/domain/model"
)

type TagRepository interface {
	Exists(model.TagName) (bool, error)
	FindByTagName(model.TagName) (bool, *model.Tag, error)
	Register(*model.Tag) error
}