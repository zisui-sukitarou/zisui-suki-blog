package model

import "time"

type Tag struct {
	TagName   TagName
	Icon      TagIcon
	Status    TagStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

/* constructor */
func NewTag(
	tagName TagName,
	icon TagIcon,
	status TagStatus,
	createdAt time.Time,
	updatedAt time.Time,
) *Tag {
	return &Tag{
		TagName:   tagName,
		Icon:      icon,
		Status:    status,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
