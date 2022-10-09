package model

import "time"

type Like struct {
	BlogId    BlogId
	UserId    UserId
	CreatedAt time.Time
}

/* constructor */
func NewLike(
	blogId BlogId,
	userId UserId,
) *Like {
	return &Like{
		BlogId:    blogId,
		UserId:    userId,
		CreatedAt: time.Now(),
	}
}
