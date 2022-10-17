package repository

import (
	"time"
	"zisui-suki-blog/domain/model"
)

/* user repository */
// * PK で Find するメソッドは、返り値の1つ目が存在フラグ
type UserRepository interface {
	FindById(model.UserId) (bool, *UserData, error)
	FindByEmail(model.UserEmail) (bool, *UserData, error)
	Register(*model.User) error
	Update(*model.User) (*UserData, error)
	Delete(*model.User) error
}

/*** full data ***/
type UserData struct {
	UserId    string    `json:"user_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Icon      string    `json:"icon"`
	CreatedAt time.Time `json:"created_at"`
}

func NewUserData(
	userId string,
	name string,
	email string,
	icon string,
	createdAt time.Time,
) *UserData {
	return &UserData{
		UserId:    userId,
		Name:      name,
		Email:     email,
		Icon:      icon,
		CreatedAt: createdAt,
	}
}
