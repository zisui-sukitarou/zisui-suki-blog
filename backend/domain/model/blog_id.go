package model

type BlogId int

func NewBlogId(id int) (BlogId, error){
	return BlogId(id), nil
}