package gateway_test

import (
	"context"
	"testing"
	"time"
	"zisui-suki-blog/adapter/gateway"
	"zisui-suki-blog/domain/model"
	"zisui-suki-blog/domain/repository"
	"zisui-suki-blog/domain/service"
	"zisui-suki-blog/infrastructure/db"
)

type userRegisterData struct {
	name        string
	displayName string
	email       string
	password    string
	icon        string
	status      uint
}

func register(u *userRegisterData, repo repository.UserRepository) error {
	s := service.User{}

	/* input data -> model */
	name, err := model.NewUserName(u.name)
	if err != nil {
		return err
	}

	displayName, err := model.NewUserDisplayName(u.displayName)
	if err != nil {
		return err
	}

	email, err := model.NewUserEmail(u.email)
	if err != nil {
		return err
	}

	password, err := model.NewUserRawPassword(u.password)
	if err != nil {
		return err
	}

	icon, err := model.NewUserIcon(u.icon)
	if err != nil {
		return err
	}

	status, err := model.NewUserStatus(u.status)
	if err != nil {
		return err
	}

	/* */
	id, err := s.GenULID()
	if err != nil {
		return err
	}

	hash, err := s.HashPassword(password)
	if err != nil {
		return err
	}

	/* */
	user := model.NewUser(
		id,
		name,
		displayName,
		email,
		hash,
		icon,
		status,
		time.Now(),
		time.Now(),
	)

	return repo.Register(user)
}

func TestRegister(t *testing.T) {

	data := []*userRegisterData{
		{
			name: "zisui-sukitarou",
			displayName: "XISUI!",
			email: "xxx@gmail.com",
			password: "passw0rd",
			icon: "yyy.png",
			status: 0,
		},
	}

	db, err := db.Init()
	if err != nil {
		t.Errorf(err.Error())
	}
	ctx := context.Background()
	repo := gateway.NewUserGateway(db, &ctx)

	for _, d := range data {
		if err := register(d, repo); err != nil {
			t.Errorf(err.Error())
		}
	}
}
