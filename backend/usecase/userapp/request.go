package userapp

import "cook-blog/domain/model"

type UserLoginRequest struct {
	Name model.UserName `json:"name"`
	Password model.UserPassword `json:"password"`
}

type UserLogoutRequest struct {
	Name model.
}