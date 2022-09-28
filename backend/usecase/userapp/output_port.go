package userapp

import "cook-blog/domain/model"

type UserOutputPort interface {
	RespondUser(UserResponse)
	RespondLoginSuccess(UserResponse) // user_id を元に JWT token を作成
	RespondLogoutSuccess(model.UserId)
	RespondSignupSuccess(UserResponse)
	RespondSignupFailure()
	RespondAuthenticationFailure()
	RespondErorr(error)
}
