package gateway_test

import (
	"context"
	"errors"
	"log"
	"testing"
	"zisui-suki-blog/adapter/gateway"
	"zisui-suki-blog/domain/model"
	"zisui-suki-blog/domain/repository"
	"zisui-suki-blog/infrastructure/db"
)

type userFindByNameData struct {
	name string
}

func findByName(u *userFindByNameData, repo repository.UserRepository) (bool, *model.User, error) {

	/* input data -> model */
	name, err := model.NewUserName(u.name)
	if err != nil {
		return false, &model.User{}, err
	}

	return repo.FindByName(name)
}

func TestFindByName(t *testing.T) {

	okData := []*userFindByNameData{
		{
			name: "zisui-sukitarou",
		},
	}

	ngData := []*userFindByNameData{
		{
			name: "doesnot-exists",
		},
	}

	db, err := db.Init()
	if err != nil {
		t.Errorf(err.Error())
	}
	ctx := context.Background()
	repo := gateway.NewUserGateway(db, &ctx)

	for _, d := range okData {
		exists, user, err := findByName(d, repo) 
		if err != nil {
			t.Errorf(err.Error())
		}
		if !exists {
			t.Error(errors.New("not eixtsts"))
		}
		log.Println("user:", user)
	}

	for _, d := range ngData {
		exists, user, _ := findByName(d, repo) 
		if exists {
			t.Error(errors.New("should not exists"))
		}
		log.Println("user:", user)
	}
}
