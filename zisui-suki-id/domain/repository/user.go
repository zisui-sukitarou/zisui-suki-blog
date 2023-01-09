package repository

import (
	"zisui-suki-blog/domain/model"
)

/* user repository */
// * Unique Key で Find するメソッドは、返り値の1つ目が存在フラグ
type UserRepository interface {
	FindById(model.UserId) (bool, *model.User, error)
	FindByName(model.UserName) (bool, *model.User, error)
	Register(*model.User) error
	Update(*model.User) (*model.User, error)
	Delete(*model.User) error
}
