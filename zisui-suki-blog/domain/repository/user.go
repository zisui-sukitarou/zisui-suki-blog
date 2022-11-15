package repository

import (
	"zisui-suki-blog/domain/model"
)

/* user repository */
// * PK で Find するメソッドは、返り値の1つ目が存在フラグ
type UserRepository interface {
	FindById(model.UserId) (bool, *model.User, error)
	FindByUserName(model.UserName) (bool, *model.User, error)
}
