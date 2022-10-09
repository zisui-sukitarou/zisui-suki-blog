package userapp

type UserLoginRequest struct {
	UserId      string `json:"user_id"`
	RawPassword string `json:"password"`
}

type UserLogoutRequest struct {
	UserId string `json:"user_id"`
}

type UserSignUpRequest struct {
	UserId      string `json:"user_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	RawPassword string `json:"password"`
	Icon        string `json:"icon"`
}

type UserFindByIdRequest struct {
	UserId string `json:"user_id"`
}

type UserUpdateRequest struct {
	UserId string `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Icon   string `json:"icon"`
}

type UserUpdatePasswordRequest struct {
	UserId      string `json:"user_id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type UserDeleteRequest struct {
	UserId      string `json:"user_id"`
	RawPassword string `json:"password"`
}
