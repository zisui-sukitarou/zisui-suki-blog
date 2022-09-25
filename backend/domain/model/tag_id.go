package model

type TagId int

func NewTagId(tagId int) (TagId, error) {
	return TagId(tagId), nil
}