package userapp

import "zisui-suki-blog/usecase/apperr"

type UserInputPort interface {
	Login(*UserLoginRequest) error
	SignUp(*UserSignUpRequest) error
	FindById(*UserFindByIdRequest) error
	FindByName(*UserFindByNameRequest) error
	Update(*UserUpdateRequest) error
	UpdatePassword(*UserUpdatePasswordRequest) error
	Delete(*UserDeleteRequest) error
}

type UserOutputPort interface {
	Login(*LoginResponse) error // user_id を元に JWT token を作成
	SignUp(*SignUpResponse) error
	FindById(*UserResponse) error // blog api から使われる
	FindByName(*UserResponse) error
	Update(*UpdateResponse) error
	UpdatePassword(*UpdatePasswordResponse) error
	Delete() error
	/* err */
	RespondError(*apperr.ErrorResponse) error
}