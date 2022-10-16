package userapp

import "zisui-suki-blog/usecase/apperr"

type UserInputPort interface {
	Login(*UserLoginRequest) error
	Logout(*UserLogoutRequest) error
	SignUp(*UserSignUpRequest) error
	FindById(*UserFindByIdRequest) error
	Update(*UserUpdateRequest) error
	UpdatePassword(*UserUpdatePasswordRequest) error
	Delete(*UserDeleteRequest) error
}

type UserOutputPort interface {
	RespondUser(*UserResponse) error
	RespondLoginSuccess(*UserResponse) error // user_id を元に JWT token を作成
	RespondLogoutSuccess(*UserResponse) error
	RespondSignupSuccess(*UserResponse) error
	RespondSignupFailure() error
	RespondAuthenticationFailure() error
	RespondError(*apperr.ErrorResponse) error
}