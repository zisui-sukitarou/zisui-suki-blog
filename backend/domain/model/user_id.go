package model

type UserId int

func NewUserId(userId int) (UserId, error) {
	return UserId(userId), nil
}
