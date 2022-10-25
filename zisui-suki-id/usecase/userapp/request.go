package userapp

type UserLoginRequest struct {
	Name        string `json:"name"`
	RawPassword string `json:"password"`
}

type UserSignUpRequest struct {
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	RawPassword string `json:"password"`
	Icon        string `json:"icon"`
	Status      uint   `json:"status"`
}

type UserFindByIdRequest struct {
	UserId string `json:"user_id"`
}

type UserFindByNameRequest struct {
	Name string `json:"name"`
}

type UserUpdateRequest struct {
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	Icon        string `json:"icon"`
}

type UserUpdatePasswordRequest struct {
	Name        string `json:"name"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type UserDeleteRequest struct {
	Name        string `json:"name"`
	RawPassword string `json:"password"`
}
