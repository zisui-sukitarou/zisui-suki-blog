package blogapp

import (
	"time"
	"zisui-suki-blog/domain/model"
	"zisui-suki-blog/domain/repository"
)

type WriterInfo struct {
	UserId      string `json:"user_id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	Icon        string `json:"icon"`
}

type TagInfo struct {
	TagName string `json:"tags_name"`
	Icon    string `json:"icon"`
}

/*** Full Response ***/
type BlogResponse struct {
	BlogId     string      `json:"blog_id"`
	Content    string      `json:"content"`
	Title      string      `json:"title"`
	Abstract   string      `json:"abstract"`
	Evaluation uint        `json:"evaluation"`
	Status     uint        `json:"status"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	Writer     *WriterInfo `json:"writer"`
	Tags       []*TagInfo  `json:"tags"`
}

func NewBlogResponse(b *repository.BlogData, u *model.User) *BlogResponse {
	writerInfo := &WriterInfo{
		UserId:      string(u.UserId),
		Name:        string(u.Name),
		DisplayName: string(u.DisplayName),
		Email:       string(u.Email),
		Icon:        string(u.Icon),
	}

	var tags []*TagInfo
	for _, tag := range b.Tags {
		tags = append(tags, &TagInfo{
			TagName: string(tag.TagName),
			Icon:    string(tag.Icon),
		})
	}

	return &BlogResponse{
		BlogId:     string(b.Blog.BlogId),
		Content:    string(b.Blog.Content),
		Title:      string(b.Blog.Title),
		Abstract:   string(b.Blog.Abstract),
		Evaluation: uint(b.Blog.Evaluation),
		Status:     uint(b.Blog.Status),
		CreatedAt:  b.Blog.CreatedAt,
		UpdatedAt:  b.Blog.UpdatedAt,
		Writer:     writerInfo,
		Tags:       tags,
	}
}

/*** BlogReponse without Content ***/
type BlogOverviewResponse struct {
	BlogId     string      `json:"blog_id"`
	Title      string      `json:"title"`
	Abstract   string      `json:"abstract"`
	Evaluation uint        `json:"evaluation"`
	Status     uint        `json:"status"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	Writer     *WriterInfo `json:"writer"`
	Tags       []*TagInfo  `json:"tags"`
}

func NewBlogOverviewResponse(b *repository.BlogData, u *model.User) *BlogOverviewResponse {
	writerInfo := &WriterInfo{
		UserId:      string(u.UserId),
		Name:        string(u.Name),
		DisplayName: string(u.DisplayName),
		Email:       string(u.Email),
		Icon:        string(u.Icon),
	}

	var tags []*TagInfo
	for _, tag := range b.Tags {
		tags = append(tags, &TagInfo{
			TagName: string(tag.TagName),
			Icon:    string(tag.Icon),
		})
	}

	return &BlogOverviewResponse{
		BlogId:     string(b.Blog.BlogId),
		Title:      string(b.Blog.Title),
		Abstract:   string(b.Blog.Abstract),
		Evaluation: uint(b.Blog.Evaluation),
		Status:     uint(b.Blog.Status),
		CreatedAt:  b.Blog.CreatedAt,
		UpdatedAt:  b.Blog.UpdatedAt,
		Writer:     writerInfo,
		Tags:       tags,
	}
}
