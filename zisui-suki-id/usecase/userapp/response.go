package userapp

import (
	"time"
	"zisui-suki-blog/domain/model"
)

/* return user info */
type UserResponse struct {
	Name        string    `json:"name"`
	DisplayName string    `json:"displayName"`
	Email       string    `json:"email"`
	Icon        string    `json:"icon"`
	Status      uint      `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewUserResponse(u *model.User) *UserResponse {
	return &UserResponse{
		Name:        string(u.Name),
		DisplayName: string(u.DisplayName),
		Email:       string(u.Email),
		Icon:        string(u.Icon),
		Status:      uint(u.Status),
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
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

/* update password */
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

/* name -> id */
type UserNameToUserIdResponse struct {
	UserId string `json:"userId"`
}

func NewUserNameToUserIdResponse(
	userId model.UserId,
) *UserNameToUserIdResponse {
	return &UserNameToUserIdResponse{
		UserId: string(userId),
	}
}