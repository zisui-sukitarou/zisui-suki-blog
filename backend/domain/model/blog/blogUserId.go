package model

type BlogUserId int

func NewBlogUserId(userId int) (BlogUserId, error){
	return BlogUserId(userId), nil
}