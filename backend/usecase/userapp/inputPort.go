package userapp

type UserInputPort interface {
	Login(UserLoginRequest)
	Logout(UserLogoutRequest)
	SignUp(UserSignUpRequest)
	FindById(UserFindByIdRequest)
	UpdateName(UserUpdateNameRequest)
	UpdateEmail(UserUpdateEmailRequest)
	UpdatePassword(UserUpdatePasswordRequest)
	Delete(UserDeleteRequest)
}
