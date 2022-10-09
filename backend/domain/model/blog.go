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
) *Blog {
	return &Blog{
		BlogId:     blogId,
		UserId:     userId,
		Content:    content,
		Title:      title,
		Abstract:   abstract,
		Evaluation: evaluation,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
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
