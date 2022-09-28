package userapp

import (
	"cook-blog/domain/model"
	"cook-blog/domain/repository"
	"cook-blog/domain/service"
	"errors"
)

type UserInteractor struct {
	OutputPort UserOutputPort
	Repository repository.UserRepository
}

/* constructor */
func NewUserInteractor(
	outputPort UserOutputPort,
	repo repository.UserRepository,
) UserInputPort {
	return &UserInteractor{
		OutputPort: outputPort,
		Repository: repo,
	}
}

func (u *UserInteractor) Login(request UserLoginRequest) {
	/* input data -> model object */
	userId, err := model.NewUserId(request.UserId)
	if err != nil {
		u.OutputPort.RespondErorr(err)
	}

	rawPassword, err := model.NewUserRawPassword(request.RawPassword)
	if err != nil {
		u.OutputPort.RespondErorr(err)
	}

	/* login */
	userService := service.User{}
	exists, user, err := u.Repository.FindById(userId)
	if !exists {
		u.OutputPort.RespondAuthenticationFailure()
	}
	if err != nil {
		u.OutputPort.RespondErorr(err)
	}

	ok, err := userService.IsConfigured(user.HashedPassword, rawPassword)
	if err != nil {
		u.OutputPort.RespondErorr(err)
	}

	if ok {
		u.OutputPort.RespondLoginSuccess(NewUserResponse(user))
	} else {
		u.OutputPort.RespondAuthenticationFailure()
	}
}

func (u *UserInteractor) Logout(request UserLogoutRequest) {
	/* input data -> model object */
	userId, err := model.NewUserId(request.UserId)
	if err != nil {
		u.OutputPort.RespondErorr(err)
	}

	/* logout */
	u.OutputPort.RespondLogoutSuccess(userId)
}

func (u *UserInteractor) SignUp(request UserSignUpRequest) {
	/* input data -> model object */
	userId, err := model.NewUserId(request.UserId)
	if err != nil {
		u.OutputPort.RespondErorr(err)
	}

	name, err := model.NewUserName(request.Name)
	if err != nil {
		u.OutputPort.RespondErorr(err)
	}

	email, err := model.NewUserEmail(request.Email)
	if err != nil {
		u.OutputPort.RespondErorr(err)
	}

	rawPassword, err := model.NewUserRawPassword(request.RawPassword)
	if err != nil {
		u.OutputPort.RespondErorr(err)
	}

	/* hash password */
	userService := service.User{}
	hashedPassword, err := userService.HashPassword(rawPassword)
	if err != nil {
		u.OutputPort.RespondErorr(err)
	}

	/* create new user object */
	command := model.CommandNewUser{
		UserId: userId,
		Name: name,
		Email: email,
		HashedPassword: hashedPassword,
	}
	user, err := model.NewUser(command)
	if err != nil {
		u.OutputPort.RespondErorr(err)
	}

	/* check overlapping (user id) */
	exits, err := userService.Exists(user.UserId, u.Repository)
	if err != nil {
		u.OutputPort.RespondErorr(err)
	}
	if exits {
		u.OutputPort.RespondSignupFailure()
	}

	/* save */
	u.Repository.Register(user)

	/* response */
	u.OutputPort.RespondSignupSuccess(NewUserResponse(user))
}

func (u *UserInteractor) FindById(request UserFindByIdRequest) {
	/* input data -> model object */
	userId, err := model.NewUserId(request.UserId)
	if err != nil {
		u.OutputPort.RespondErorr(err)
	}

	/* find user by id */
	exists, user, err := u.Repository.FindById(userId)
	if err != nil {
		u.OutputPort.RespondErorr(err)
	}
	if !exists {
		u.OutputPort.RespondErorr(errors.New("user of the id dosen't exists"))
	}

	/* response */
	u.OutputPort.RespondUser(NewUserResponse(user))
}

func (u *UserInteractor) Update(request UserUpdateRequest) {
	/* input data -> model object */
	userId, err := model.NewUserId(request.UserId)
	if err != nil {
		u.OutputPort.RespondErorr(err)
	}

	name, err := model.NewUserName(request.Name)
	if err != nil {
		u.OutputPort.RespondErorr(err)
	}

	email, err := model.NewUserEmail(request.Email)
	if err != nil {
		u.OutputPort.RespondErorr(err)
	}

	/* find user by id */
	exists, user, err := u.Repository.FindById(userId)
	if err != nil {
		u.OutputPort.RespondErorr(err)
	}
	if !exists {
		u.OutputPort.RespondErorr(errors.New("user of the id dosen't exists"))
	}

	/* update & save */
	user.UpdateName(name)
	user.UpdateEmail(email)
	nUser, err := u.Repository.Update(user)
	if err != nil {
		u.OutputPort.RespondErorr(err)
	}

	/* response */
	u.OutputPort.RespondUser(NewUserResponse(nUser))
}

func (u *UserInteractor) UpdatePassword(request UserUpdatePasswordRequest) {
	/* input data -> model object */
	userId, err := model.NewUserId(request.UserId)
	if err != nil {
		u.OutputPort.RespondErorr(err)
	}

	oldPass, err := model.NewUserRawPassword(request.OldPassword)
	if err != nil {
		u.OutputPort.RespondErorr(err)
	}

	newPass, err := model.NewUserRawPassword(request.NewPassword)
	if err != nil {
		u.OutputPort.RespondErorr(err)
	}

	/* configure user */
	exists, user, err := u.Repository.FindById(userId)
	if err != nil {
		u.OutputPort.RespondErorr(err)
	}
	if !exists {
		u.OutputPort.RespondErorr(errors.New("user of the id dosen't exists"))
	}

	userService := service.User{}
	ok, err := userService.IsConfigured(user.HashedPassword, oldPass)
	if err != nil {
		u.OutputPort.RespondErorr(err)
	}
	if !ok {
		u.OutputPort.RespondAuthenticationFailure()
	}

	/* update password & save */
	newHash, err := userService.HashPassword(newPass)
	if err != nil {
		u.OutputPort.RespondErorr(err)
	}
	user.UpdatePassword(newHash)
	u.Repository.Update(user)

	/* response */
	u.OutputPort.RespondUser(NewUserResponse(user))
}

func (u *UserInteractor) Delete(request UserDeleteRequest) {
	
}
