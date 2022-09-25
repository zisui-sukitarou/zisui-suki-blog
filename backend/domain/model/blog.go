package model

import (
	"time"
)

type Blog struct {
	Id        BlogId
	UserId    UserId
	Content   BlogContent
	Title     BlogTitle
	Abstract  BlogAbstract
	View      BlogView
	CreatedAt time.Time
	UpdatedAt time.Time
}

/* constructor */
type CommandNewBlog struct {
	UserId   UserId
	Content  BlogContent
	Title    BlogTitle
	Abstract BlogAbstract
}

func NewBlog(command CommandNewBlog) (*Blog, error) {
	return &Blog{
		UserId:   command.UserId,
		Content:  command.Content,
		Title:    command.Title,
		Abstract: command.Abstract,
	}, nil
}

/* change content */
type CommandUpdateBlog struct {
	Content  BlogContent
	Title    BlogTitle
	Abstract BlogAbstract
}

func (b *Blog) Update(command CommandUpdateBlog) error {
	b.Content = command.Content
	b.Title = command.Title
	b.Abstract = command.Abstract
	return nil
}

/* increment view */
func (b *Blog) IncView() error {
	b.View++
	return nil
}
