package gateway_test

import (
	"testing"
	"time"
	"zisui-suki-blog/adapter/gateway"
	"zisui-suki-blog/domain/model"
	"zisui-suki-blog/infrastructure/db"
)

func TestTagRegister(t *testing.T) {
	tag := model.NewTag(
		model.TagName("ramen"),
		time.Now(),
	)

	infra, err := db.Init()
	if err != nil {
		t.Errorf(err.Error())
	}

	err = gateway.NewTagGateway(infra.Client, &infra.Ctx).Register(tag)
	if err != nil {
		t.Errorf(err.Error())
	}
}