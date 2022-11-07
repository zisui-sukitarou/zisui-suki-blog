package userapp

import "zisui-suki-blog/domain/model"

/* return user info */
type UserResponse struct {
	UserId      string `json:"userId"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
	Icon        string `json:"icon"`
	Status      uint   `json:"status"`
}

func NewUserResponse(u *model.User) *UserResponse {
	return &UserResponse{
		UserId:      string(u.UserId),
		Name:        string(u.Name),
		DisplayName: string(u.DisplayName),
		Email:       string(u.Email),
		Icon:        string(u.Icon),
		Status:      uint(u.Status),
	}
}

/* */
type SignUpResponse struct {
	Status   uint          `json:"status"`
	JWT      string        `json:"token"`
	UserInfo *UserResponse `json:"userInfo"`
}

func NewSignUpResponse(
	status uint,
	jwt string,
	user *model.User,
) *SignUpResponse {
	return &SignUpResponse{
		Status:   status,
		JWT:      jwt,
		UserInfo: NewUserResponse(user),
	}
}

/* */
type LoginResponse struct {
	Status   uint          `json:"status"`
	JWT      string        `json:"token"`
	UserInfo *UserResponse `json:"userInfo"`
}

func NewLoginResponse(
	status uint,
	jwt string,
	user *model.User,
) *LoginResponse {
	return &LoginResponse{
		Status:   status,
		JWT:      jwt,
		UserInfo: NewUserResponse(user),
	}
}

/* */
type UpdateResponse struct {
	Status   uint          `json:"status"`
	UserInfo *UserResponse `json:"userInfo"`
}

func NewUpdateResponse(
	status uint,
	user *model.User,
) *UpdateResponse {
	return &UpdateResponse{
		Status:   status,
		UserInfo: NewUserResponse(user),
	}
}

/* */
type UpdatePasswordResponse struct {
	Status   uint          `json:"status"`
	UserInfo *UserResponse `json:"userInfo"`
}

func NewUpdatePasswordResponse(
	status uint,
	user *model.User,
) *UpdatePasswordResponse {
	return &UpdatePasswordResponse{
		Status:   status,
		UserInfo: NewUserResponse(user),
	}
}
