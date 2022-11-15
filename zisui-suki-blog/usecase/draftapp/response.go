package draftapp

import (
	"time"
	"zisui-suki-blog/domain/model"
	"zisui-suki-blog/domain/repository"
)

type WriterInfo struct {
	UserId      string `json:"userId"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
	Icon        string `json:"icon"`
}

type TagInfo struct {
	TagName string `json:"tagsName"`
	Icon    string `json:"icon"`
}

/*** Full Response ***/
type DraftResponse struct {
	BlogId     string      `json:"blogId"`
	Content    string      `json:"content"`
	Title      string      `json:"title"`
	Abstract   string      `json:"abstract"`
	Evaluation uint        `json:"evaluation"`
	Status     uint        `json:"status"`
	CreatedAt  time.Time   `json:"createdAt"`
	UpdatedAt  time.Time   `json:"updatedAt"`
	Writer     *WriterInfo `json:"writer"`
	Tags       []*TagInfo  `json:"tags"`
}

func NewDraftResponse(d *repository.DraftData, u *model.User) *DraftResponse {
	writerInfo := &WriterInfo{
		UserId:      string(u.UserId),
		Name:        string(u.Name),
		DisplayName: string(u.DisplayName),
		Email:       string(u.Email),
		Icon:        string(u.Icon),
	}

	var tags []*TagInfo
	for _, tag := range d.Tags {
		tags = append(tags, &TagInfo{
			TagName: string(tag.TagName),
			Icon:    string(tag.Icon),
		})
	}

	return &DraftResponse{
		BlogId:     string(d.Draft.DraftId),
		Content:    string(d.Draft.Content),
		Title:      string(d.Draft.Title),
		Abstract:   string(d.Draft.Abstract),
		Evaluation: uint(d.Draft.Evaluation),
		Status:     uint(d.Draft.Status),
		CreatedAt:  d.Draft.CreatedAt,
		UpdatedAt:  d.Draft.UpdatedAt,
		Writer:     writerInfo,
		Tags:       tags,
	}
}

/*** BlogReponse without Content ***/
type DraftOverviewResponse struct {
	BlogId     string      `json:"blogId"`
	Title      string      `json:"title"`
	Abstract   string      `json:"abstract"`
	Evaluation uint        `json:"evaluation"`
	Status     uint        `json:"status"`
	CreatedAt  time.Time   `json:"createdAt"`
	UpdatedAt  time.Time   `json:"updatedAt"`
	Writer     *WriterInfo `json:"writer"`
	Tags       []*TagInfo  `json:"tags"`
}

func NewDraftOverviewResponse(d *repository.DraftData, u *model.User) *DraftOverviewResponse {
	writerInfo := &WriterInfo{
		UserId:      string(u.UserId),
		Name:        string(u.Name),
		DisplayName: string(u.DisplayName),
		Email:       string(u.Email),
		Icon:        string(u.Icon),
	}

	var tags []*TagInfo
	for _, tag := range d.Tags {
		tags = append(tags, &TagInfo{
			TagName: string(tag.TagName),
			Icon:    string(tag.Icon),
		})
	}

	return &DraftOverviewResponse{
		BlogId:     string(d.Draft.DraftId),
		Title:      string(d.Draft.Title),
		Abstract:   string(d.Draft.Abstract),
		Evaluation: uint(d.Draft.Evaluation),
		Status:     uint(d.Draft.Status),
		CreatedAt:  d.Draft.CreatedAt,
		UpdatedAt:  d.Draft.UpdatedAt,
		Writer:     writerInfo,
		Tags:       tags,
	}
}
