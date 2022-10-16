package blogapp

import (
	"time"
	"zisui-suki-blog/domain/model"
)

/*** Full Response ***/
type BlogResponse struct {
	BlogId    string    `json:"blog_id"`
	UserId    string    `json:"user_id"`
	Content   string    `json:"content"`
	Title     string    `json:"title"`
	Abstract  string    `json:"abstract"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewBlogResponse(b *model.Blog) *BlogResponse {
	return &BlogResponse{
		BlogId:    string(b.BlogId),
		UserId:    string(b.UserId),
		Content:   string(b.Content),
		Title:     string(b.Title),
		Abstract:  string(b.Abstract),
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
	}
}

/*** BlogReponse without Content ***/
type BlogOverviewResponse struct {
	BlogId    string    `json:"blog_id"`
	UserId    string    `json:"user_id"`
	Title     string    `json:"title"`
	Abstract  string    `json:"abstract"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewBlogOverviewResponse(b *model.Blog) *BlogOverviewResponse {
	return &BlogOverviewResponse{
		BlogId:    string(b.BlogId),
		UserId:    string(b.UserId),
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
