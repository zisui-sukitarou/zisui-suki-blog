package model

import (
	"time"
)

type Blog struct {
	BlogId     BlogId
	UserId     UserId
	Content    BlogContent
	Title      BlogTitle
	Abstract   BlogAbstract
	Evaluation BlogEvaluation
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

/* constructor */
func NewBlog(
	blogId BlogId,
	userId UserId,
	content BlogContent,
	title BlogTitle,
	abstract BlogAbstract,
	evaluation BlogEvaluation,
	createdAt time.Time,
	updatedAt time.Time,
) *Blog {
	return &Blog{
		BlogId:     blogId,
		UserId:     userId,
		Content:    content,
		Title:      title,
		Abstract:   abstract,
		Evaluation: evaluation,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
	}
}

/* change content */
func (b *Blog) Update(
	content BlogContent,
	title BlogTitle,
	abstract BlogAbstract,
	evaluation BlogEvaluation,
) {
	b.Content = content
	b.Title = title
	b.Abstract = abstract
	b.Evaluation = evaluation
	b.UpdatedAt = time.Now()
}
