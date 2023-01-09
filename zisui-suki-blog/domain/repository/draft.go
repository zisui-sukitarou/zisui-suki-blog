package repository

import (
	"zisui-suki-blog/domain/model"
)

/*** blog repository ***
* Unique Key で Find するメソッドは、返り値の1つ目が存在フラグ
* Blog 配列が返ってくるメソッドの Blog.Content は空文字列（BlogOverviewResponse に不要なため）
 */
type DraftRepository interface {
	FindById(model.DraftId) (bool, *DraftData, error)
	FindByUserId(userId model.UserId, begin uint, end uint) ([]*DraftData, error) // Blog.Content は空文字列
	Register(*model.Draft) error
	Update(*model.Draft) error
	UpdateTags(model.DraftId, []model.TagName) error
	Delete(*model.Draft) error
}

/*** full data ***
* writer の情報はこのデータベースにはない
 */
type DraftData struct {
	Draft *model.Draft
	Tags  []*model.Tag
}

func NewDraftData(
	draft *model.Draft,
	tags []*model.Tag,
) *DraftData {
	return &DraftData{
		Draft: draft,
		Tags:  tags,
	}
}
