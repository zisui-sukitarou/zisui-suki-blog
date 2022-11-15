package model

import (
	"time"
)

type Draft struct {
	DraftId    DraftId
	UserId     UserId
	Content    DraftContent
	Title      DraftTitle
	Abstract   DraftAbstract
	Evaluation DraftEvaluation
	Status     DraftStatus
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

/* constructor */
func NewDraft(
	draftId DraftId,
	userId UserId,
	content DraftContent,
	title DraftTitle,
	abstract DraftAbstract,
	evaluation DraftEvaluation,
	status DraftStatus,
	createdAt time.Time,
	updatedAt time.Time,
) *Draft {
	return &Draft{
		DraftId:    draftId,
		UserId:     userId,
		Content:    content,
		Title:      title,
		Abstract:   abstract,
		Evaluation: evaluation,
		Status:     status,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
	}
}

/* change content */
func (d *Draft) Update(
	content DraftContent,
	title DraftTitle,
	abstract DraftAbstract,
	evaluation DraftEvaluation,
) {
	d.Content = content
	d.Title = title
	d.Abstract = abstract
	d.Evaluation = evaluation
	d.UpdatedAt = time.Now()
}
