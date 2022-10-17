package userapp

import "zisui-suki-blog/domain/model"

/* return user info */
type UserResponse struct {
	UserId string `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Icon   string `json:"icon"`
}

func NewUserResponse(u *model.User) *UserResponse {
	return &UserResponse{
		UserId: string(u.UserId),
		Name:   string(u.Name),
		Email:  string(u.Email),
		Icon:   string(u.Icon),
	}
}
