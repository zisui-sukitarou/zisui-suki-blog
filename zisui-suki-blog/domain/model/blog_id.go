package model

type BlogId string

func NewBlogId(id string) (BlogId, error) {
	return BlogId(id), nil
}
