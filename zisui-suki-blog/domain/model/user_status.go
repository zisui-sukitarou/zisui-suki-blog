package model

type UserStatus uint

func NewUserStatus(status uint) (UserStatus, error) {
	return UserStatus(status), nil
}