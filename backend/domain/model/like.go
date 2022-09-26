package model

import "time"

type Like struct {
	BlogId    BlogId
	UserId    UserId
	CreatedAt time.Time
}

/* constructor */
type CommandNewLike struct {
	BlogId BlogId
	UserId UserId
}

func NewLike(command CommandNewLike) (*Like, error) {
	return &Like{
		BlogId: command.BlogId,
		UserId: command.UserId,
	}, nil
}
