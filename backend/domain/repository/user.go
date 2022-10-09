package repository

import "cook-blog/domain/model"

/* user repository */
// * PK で Find するメソッドは、返り値の1つ目が存在フラグ
type UserRepository interface {
	FindById(model.UserId) (bool, *model.User, error)
	FindByEmail(model.UserEmail) (bool, *model.User)
	Register(*model.User) error
	Update(*model.User) (*model.User, error)
	Delete(*model.User) error
}
