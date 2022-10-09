package userapp

import "cook-blog/domain/model"

type UserInputPort interface {
	Login(UserLoginRequest)
	Logout(UserLogoutRequest)
	SignUp(UserSignUpRequest)
	FindById(UserFindByIdRequest)
	Update(UserUpdateRequest)
	UpdatePassword(UserUpdatePasswordRequest)
	Delete(UserDeleteRequest)
}

type UserOutputPort interface {
	RespondUser(UserResponse)
	RespondLoginSuccess(UserResponse) // user_id を元に JWT token を作成
	RespondLogoutSuccess(model.UserId)
	RespondSignupSuccess(UserResponse)
	RespondSignupFailure()
	RespondAuthenticationFailure()
	RespondErorr(error)
}