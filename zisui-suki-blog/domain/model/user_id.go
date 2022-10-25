package model

type UserId string

func NewUserId(id string) (UserId, error) {
	return UserId(id), nil
}
