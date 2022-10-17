package blogapp

import (
	"time"
	"zisui-suki-blog/domain/model"
)

/*** Full Response ***/
type BlogResponse struct {
	BlogId    string     `json:"blog_id"`
	Writer    WriterInfo `json:"writer"`
	Content   string     `json:"content"`
	Title     string     `json:"title"`
	Abstract  string     `json:"abstract"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func NewBlogResponse(b *model.Blog, u *model.User) *BlogResponse {
	return &BlogResponse{
		BlogId:    string(b.BlogId),
		Writer:    NewWriterInfo(u),
		Content:   string(b.Content),
		Title:     string(b.Title),
		Abstract:  string(b.Abstract),
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
	}
}

/*** BlogReponse without Content ***/
type BlogOverviewResponse struct {
	BlogId    string     `json:"blog_id"`
	Writer    WriterInfo `json:"writer"`
	Title     string     `json:"title"`
	Abstract  string     `json:"abstract"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func NewBlogOverviewResponse(b *model.Blog, u *model.User) *BlogOverviewResponse {
	return &BlogOverviewResponse{
		BlogId:    string(b.BlogId),
		Writer:    NewWriterInfo(u),
		Title:     string(b.Title),
		Abstract:  string(b.Abstract),
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
	}
}

func NewBlogOverviewsResponse(bs []*model.Blog) []*BlogOverviewResponse {
	var res []*BlogOverviewResponse
	for _, v := range bs {
		res = append(res, NewBlogOverviewResponse(v))
	}
	return res
}

/*** Writer Info ***/
type WriterInfo struct {
	UserId string `user_id`
	Name   string `name`
	Icon   string `icon`
}

func NewWriterInfo(u *model.User) WriterInfo {
	return WriterInfo{
		UserId: string(u.UserId),
		Name:   string(u.Name),
		Icon:   string(u.Icon),
	}
}
