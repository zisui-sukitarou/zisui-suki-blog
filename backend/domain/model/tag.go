package model

import "time"

type Tag struct {
	Id        TagId
	Name      TagName
	CreatedAt time.Time
}

/* constructor */
type CommandNewTag struct {
	Name TagName
}

func NewTag(command CommandNewTag) (*Tag, error) {
	return &Tag{
		Name: command.Name,
	}, nil
}
