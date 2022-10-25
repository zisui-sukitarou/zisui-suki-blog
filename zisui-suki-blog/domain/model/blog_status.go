package model

type BlogStatus uint

func NewBlogStatus(status uint) (BlogStatus, error) {
	return BlogStatus(status), nil
}