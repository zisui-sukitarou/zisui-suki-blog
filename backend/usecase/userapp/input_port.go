package userapp

type UserInputPort interface {
	Login(UserLoginRequest)
	Logout(UserLogoutRequest)
	SignUp(UserSignUpRequest)
	FindById(UserFindByIdRequest)
	Update(UserUpdateRequest)
	UpdatePassword(UserUpdatePasswordRequest)
	Delete(UserDeleteRequest)
}
