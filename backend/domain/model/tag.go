package model

import "time"

type Tag struct {
	TagName      TagName
	CreatedAt time.Time
}

/* constructor */
type CommandNewTag struct {
	TagName TagName
}

func NewTag(command CommandNewTag) (*Tag, error) {
	return &Tag{
		TagName: command.TagName,
	}, nil
}
