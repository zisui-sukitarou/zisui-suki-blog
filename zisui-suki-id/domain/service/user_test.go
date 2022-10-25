package service_test

import (
	"log"
	"testing"
	"zisui-suki-blog/domain/model"
	"zisui-suki-blog/domain/service"
)

func TestGenJWT(t *testing.T) {
	uService := service.User{}
	userId, _ := model.NewUserId("itmchineughlx")
	token, err := uService.GenJWT(userId)
	if err != nil {
		t.Errorf(err.Error())
	}
	log.Println(token)
}