package gateway_test

import (
	"testing"
	"time"
	"zisui-suki-blog/adapter/gateway"
	"zisui-suki-blog/domain/model"
	"zisui-suki-blog/domain/service"
	"zisui-suki-blog/infrastructure/db"
)

func TestUserRegister(t *testing.T) {

	hashed, err := newUserService().HashPassword("password")
	if err != nil {
		t.Errorf(err.Error())
	}

	user := model.NewUser(
		model.UserId("zisui-sukitarou"),
		model.UserName("ZISUI!"),
		model.UserEmail("xxx@gmail.com"),
		model.UserHashedPassword(hashed),
		model.UserIcon("xxx.jpg"),
		time.Now(),
		time.Now(),
	)

	infra, err := db.Init()
	if err != nil {
		t.Errorf(err.Error())
	}

	gateway.NewUserGateway(infra.Client, &infra.Ctx).Register(user)
}

func newUserService() *service.User {
	return &service.User{}
}
