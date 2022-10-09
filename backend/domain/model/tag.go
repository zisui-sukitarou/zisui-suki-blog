package model

import "time"

type Tag struct {
	TagName   TagName
	CreatedAt time.Time
}

/* constructor */
func NewTag(tagName TagName) *Tag {
	return &Tag{
		TagName:   tagName,
		CreatedAt: time.Now(),
	}
}
