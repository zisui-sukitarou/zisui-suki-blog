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

func (u *UserInteractor) NameToUserId(request *UserNameToUserIdRequest) error {
	/* input data -> mode object */
	name, err := model.NewUserName(request.Name)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}

	/* does user exist ? */
	status := Status{}
	exists, user, err := u.Repository.FindByName(name)
	if !exists {
		return u.OutputPort.Login(NewLoginResponse(
			status.UserNotExists(),
			"",
			&model.User{},
		))
	}
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}

	/* response */
	return u.OutputPort.NameToUserId(NewUserNameToUserIdResponse(user.UserId))
}

func (u *UserInteractor) Login(request *UserLoginRequest) error {
	/* input data -> model object */
	name, err := model.NewUserName(request.Name)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}

	rawPassword, err := model.NewUserRawPassword(request.RawPassword)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}

	/* does user exist ? */
	status := Status{}
	exists, user, err := u.Repository.FindByName(name)
	if !exists {
		return u.OutputPort.Login(NewLoginResponse(
			status.UserNotExists(),
			"",
			&model.User{},
		))
	}
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}

	/* is password correct ? */
	ok, err := u.service().IsConfigured(user.Password, rawPassword)
	if err != nil {
		return u.OutputPort.Login(NewLoginResponse(
			status.InvalidPassword(),
			"",
			&model.User{},
		))
	}
	if !ok {
		return u.OutputPort.Login(NewLoginResponse(
			status.InvalidPassword(),
			"",
			&model.User{},
		))
	}

	/* create JWT */
	jwt, err := u.service().GenJWT(user.UserId)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}

	/* response */
	return u.OutputPort.Login(NewLoginResponse(
		status.OK(),
		jwt,
		user,
	))
}

func (u *UserInteractor) SignUp(request *UserSignUpRequest) error {
	/* input data -> model object */
	name, err := model.NewUserName(request.Name)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}

	displayName, err := model.NewUserDisplayName(request.Name)
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

	status, err := model.NewUserStatus(request.Status)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}

	/* assign id */
	userId, err := u.service().GenULID()
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}

	/* hash password */
	hashedPassword, err := u.service().HashPassword(rawPassword)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}

	/* create new user object */
	user := model.NewUser(userId, name, displayName, email, hashedPassword, icon, status, time.Now(), time.Now())

	/* check overlapping (user name) */
	resStatus := Status{}
	exits, err := u.service().Exists(user.Name, u.Repository)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}
	if exits {
		return u.OutputPort.SignUp(NewSignUpResponse(
			resStatus.UserAlreadyExists(),
			"",
			user,
		))
	}

	/* save */
	err = u.Repository.Register(user)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}

	/* generate jwt */
	jwt, err := u.service().GenJWT(userId)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}

	/* response */
	return u.OutputPort.SignUp(NewSignUpResponse(
		resStatus.OK(),
		jwt,
		user,
	))
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
	return u.OutputPort.FindById(NewUserResponse(user))
}

func (u *UserInteractor) FindByName(request *UserFindByNameRequest) error {
	/* input data -> model object */
	name, err := model.NewUserName(request.Name)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}

	/* find user by id */
	exists, user, err := u.Repository.FindByName(name)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}
	if !exists {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(errors.New("user of the name dosen't exists")))
	}

	/* response */
	return u.OutputPort.FindByName(NewUserResponse(user))
}

func (u *UserInteractor) Update(request *UserUpdateRequest) error {
	/* input data -> model object */
	name, err := model.NewUserName(request.Name)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}

	displayName, err := model.NewUserDisplayName(request.DisplayName)
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
	status := Status{}
	exists, user, err := u.Repository.FindByName(name)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}
	if !exists {
		return u.OutputPort.Update(NewUpdateResponse(
			status.UserAlreadyExists(),
			user,
		))
	}

	/* update & save */
	user.UpdateName(name)
	user.UpdateDisplayName(displayName)
	user.UpdateEmail(email)
	user.UpdateIcon(icon)
	nUser, err := u.Repository.Update(user)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}

	/* response */
	return u.OutputPort.Update(NewUpdateResponse(
		status.OK(),
		nUser,
	))
}

func (u *UserInteractor) UpdatePassword(request *UserUpdatePasswordRequest) error {
	/* input data -> model object */
	name, err := model.NewUserName(request.Name)
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
	status := Status{}
	exists, user, err := u.Repository.FindByName(name)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}
	if !exists {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(errors.New("user of the id dosen't exists")))
	}

	userService := service.User{}
	ok, err := userService.IsConfigured(user.Password, oldPass)
	if err != nil {
		return u.OutputPort.RespondError(apperr.NewErrorResponse(err))
	}
	if !ok {
		return u.OutputPort.UpdatePassword(NewUpdatePasswordResponse(
			status.InvalidPassword(),
			&model.User{},
		))
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
	return u.OutputPort.UpdatePassword(NewUpdatePasswordResponse(
		status.OK(),
		user,
	))
}

func (u *UserInteractor) Delete(request *UserDeleteRequest) error {
	return nil
}

/* service constructor */
func (u *UserInteractor) service() *service.User {
	return &service.User{}
}
