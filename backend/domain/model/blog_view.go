package model

type BlogView int

func NewBlogView(view int) (BlogView, error){
	return BlogView(view), nil
}