package usecase

import "cook-blog/domain/model"

/* return blog info */
type BlogResponse struct {
	Id       int    `json:"id"`
	UserId   int    `json:"user_id"`
	Content  string `json:"content"`
	Title    string `json:"title"`
	Abstract string `json:"abstract"`
	View     int    `json:"view"`
}

func NewBlogResponse(b *model.Blog) BlogResponse {
	return BlogResponse{
		Id:       int(b.Id),
		UserId:   int(b.UserId),
		Content:  string(b.Content),
		Title:    string(b.Title),
		Abstract: string(b.Abstract),
		View:     int(b.View),
	}
}

func NewBlogsResponse(bs []*model.Blog) []BlogResponse {
	var res []BlogResponse
	for _, v := range bs {
		res = append(res, NewBlogResponse(v))
	}
	return res
}
