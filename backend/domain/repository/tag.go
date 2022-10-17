package repository

import (
	"time"
	"zisui-suki-blog/domain/model"
)

type TagRepository interface {
	FindByTagName(model.TagName) (bool, *model.Tag, error)
	Register(*model.Tag) error
}

/*** full data ***/
type TagData struct {
	TagName   string    `json:"tag_name"`
	CreatedAt time.Time `json:"created_at"`
}

func NewTagData(
	tagName string,
	createdAt time.Time,
) *TagData {
	return &TagData{
		TagName:   tagName,
		CreatedAt: createdAt,
	}
}
