package blogapp

import (
	"cook-blog/domain/model"
	"time"
)

/* return blog info */
type BlogResponse struct {
	BlogId    string    `json:"blog_id"`
	UserId    string    `json:"user_id"`
	Content   string    `json:"content"`
	Title     string    `json:"title"`
	Abstract  string    `json:"abstract"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewBlogResponse(b *model.Blog) BlogResponse {
	return BlogResponse{
		BlogId:    string(b.BlogId),
		UserId:    string(b.UserId),
		Content:   string(b.Content),
		Title:     string(b.Title),
		Abstract:  string(b.Abstract),
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
	}
}

func NewBlogsResponse(bs []*model.Blog) []BlogResponse {
	var res []BlogResponse
	for _, v := range bs {
		res = append(res, NewBlogResponse(v))
	}
	return res
}
