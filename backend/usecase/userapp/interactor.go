package userapp

import (
	"errors"
	"time"
	"zisui-suki-blog/domain/model"
	"zisui-suki-blog/domain/repository"
	"zisui-suki-blog/domain/service"
	"zisui-suki-blog/usecase/apperr"
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

func (u *UserInteractor) Login(request *UserLoginRequest) error {
	/* input data -> model object */
	userId, err := model.NewUserId(request.UserId)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}

	rawPassword, err := model.NewUserRawPassword(request.RawPassword)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}

	/* login */
	userService := service.User{}
	exists, user, err := u.Repository.FindById(userId)
	if !exists {
		return u.OutputPort.RespondAuthenticationFailure()
	}
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}

	/* is correct password ? */
	ok, err := userService.IsConfigured(user.HashedPassword, rawPassword)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}
	if !ok {
		return u.OutputPort.RespondAuthenticationFailure()
	}
	return u.OutputPort.RespondLoginSuccess(NewUserResponse(user))
}

func (u *UserInteractor) Logout(request *UserLogoutRequest) error {
	/* input data -> model object */
	userId, err := model.NewUserId(request.UserId)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}

	/* find user */
	exists, user, err := u.Repository.FindById(userId)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}
	if !exists {
		return u.OutputPort.RespondAuthenticationFailure()
	}

	/* logout */
	return u.OutputPort.RespondLogoutSuccess(NewUserResponse(user))
}

func (u *UserInteractor) SignUp(request *UserSignUpRequest) error {
	/* input data -> model object */
	userId, err := model.NewUserId(request.UserId)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}

	name, err := model.NewUserName(request.Name)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}

	email, err := model.NewUserEmail(request.Email)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}

	rawPassword, err := model.NewUserRawPassword(request.RawPassword)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}

	icon, err := model.NewUserIcon(request.Icon)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}

	/* hash password */
	userService := service.User{}
	hashedPassword, err := userService.HashPassword(rawPassword)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}

	/* create new user object */
	user := model.NewUser(userId, name, email, hashedPassword, icon, time.Now(), time.Now())

	/* check overlapping (user id) */
	exits, err := userService.Exists(user.UserId, u.Repository)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}
	if exits {
		return u.OutputPort.RespondSignupFailure()
	}

	/* save */
	err = u.Repository.Register(user)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}

	/* response */
	return u.OutputPort.RespondSignupSuccess(NewUserResponse(user))
}

func (u *UserInteractor) FindById(request *UserFindByIdRequest) error {
	/* input data -> model object */
	userId, err := model.NewUserId(request.UserId)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}

	/* find user by id */
	exists, user, err := u.Repository.FindById(userId)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}
	if !exists {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(errors.New("user of the id dosen't exists")))
	}

	/* response */
	return u.OutputPort.RespondUser(NewUserResponse(user))
}

func (u *UserInteractor) Update(request *UserUpdateRequest) error {
	/* input data -> model object */
	userId, err := model.NewUserId(request.UserId)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}

	name, err := model.NewUserName(request.Name)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}

	email, err := model.NewUserEmail(request.Email)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}

	icon, err := model.NewUserIcon(request.Icon)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}

	/* find user by id */
	exists, user, err := u.Repository.FindById(userId)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}
	if !exists {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(errors.New("user of the id dosen't exists")))
	}

	/* update & save */
	user.UpdateName(name)
	user.UpdateEmail(email)
	user.UpdateIcon(icon)
	nUser, err := u.Repository.Update(user)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}

	/* response */
	return u.OutputPort.RespondUser(NewUserResponse(nUser))
}

func (u *UserInteractor) UpdatePassword(request *UserUpdatePasswordRequest) error {
	/* input data -> model object */
	userId, err := model.NewUserId(request.UserId)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}

	oldPass, err := model.NewUserRawPassword(request.OldPassword)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}

	newPass, err := model.NewUserRawPassword(request.NewPassword)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}

	/* configure user */
	exists, user, err := u.Repository.FindById(userId)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}
	if !exists {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(errors.New("user of the id dosen't exists")))
	}

	userService := service.User{}
	ok, err := userService.IsConfigured(user.HashedPassword, oldPass)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}
	if !ok {
		return u.OutputPort.RespondAuthenticationFailure()
	}

	/* update password & save */
	newHash, err := userService.HashPassword(newPass)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}
	user.UpdatePassword(newHash)
	user, err = u.Repository.Update(user)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}

	/* response */
	return u.OutputPort.RespondUser(NewUserResponse(user))
}

func (u *UserInteractor) Delete(request *UserDeleteRequest) error {
	return nil
}
