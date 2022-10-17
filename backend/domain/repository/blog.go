package repository

import (
	"time"
	"zisui-suki-blog/domain/model"
)

/*** blog repository ***
* PK で Find するメソッドは、返り値の1つ目が存在フラグ
* Blog 配列が返ってくるメソッドの Blog.Content は空文字列（BlogOverviewResponse に不要なため）
 */
type BlogRepository interface {
	FindById(model.BlogId) (bool, *BlogData, error)
	FindByUserId(userId model.UserId, begin uint, end uint) ([]*BlogOverviewData, error)                                  // Blog.Content は空文字列
	FindByTagName(tagName model.TagName, begin uint, end uint) ([]*BlogOverviewData, error)                               // Blog.Content は空文字列
	FindByUserIdAndTagName(userId model.UserId, tagName model.TagName, begin uint, end uint) ([]*BlogOverviewData, error) // Blog.Content は空文字列
	Register(*model.Blog) error
	RegisterTags(model.BlogId, []model.TagName) error
	Update(*model.Blog) (*BlogData, error)
	Delete(*model.Blog) error
}

/*** full data ***/
type BlogData struct {
	BlogId    string     `json:"blog_id"`
	Content   string     `json:"content"`
	Title     string     `json:"title"`
	Abstract  string     `json:"abstract"`
	Writer    *UserData  `json:"writer"`
	Tags      []*TagData `json:"tags"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func NewBlogData(
	blogId string,
	content string,
	title string,
	abstract string,
	writer *UserData,
	tags []*TagData,
	createdAt time.Time,
	updatedAt time.Time,
) *BlogData {
	return &BlogData{
		BlogId:    blogId,
		Content:   content,
		Title:     title,
		Abstract:  abstract,
		Writer:    writer,
		Tags:      tags,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

/*** BlogData から Content をなくしたもの ***/
type BlogOverviewData struct {
	BlogId    string     `json:"blog_id"`
	Title     string     `json:"title"`
	Abstract  string     `json:"abstract"`
	Writer    *UserData  `json:"writer"`
	Tags      []*TagData `json:"tags"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func NewBlogOverviewData(
	blogId string,
	title string,
	abstract string,
	writer *UserData,
	tags []*TagData,
	createdAt time.Time,
	updatedAt time.Time,
) *BlogOverviewData {
	return &BlogOverviewData{
		BlogId:    blogId,
		Title:     title,
		Abstract:  abstract,
		Writer:    writer,
		Tags:      tags,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
