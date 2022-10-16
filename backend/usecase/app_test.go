package usecase_test

import (
	"zisui-suki-blog/domain/model"
	"zisui-suki-blog/domain/service"
	"log"
	"testing"
)

func hashPassword(rawPassword model.UserRawPassword) (model.UserHashedPassword, error) {
	userService := service.User{}
	hashed, err := userService.HashPassword(rawPassword)
	if err != nil {
		return model.UserHashedPassword(""), err
	}

	return hashed, nil
}

func isConfigured(raw model.UserRawPassword, hash model.UserHashedPassword) (bool, error) {
	userService := service.User{}
	return userService.IsConfigured(hash, raw)
}

func TestHashPassword(t *testing.T) {
	testData := []string{
		"password",
		"p@ssw0rd",
		"Zisui_5kit4rou!",
		"hogefugapiyo",
	}

	var raws []model.UserRawPassword

	for _, v := range testData {
		raw, err := model.NewUserRawPassword(v)
		if err != nil {
			t.Errorf(err.Error())
		}
		raws = append(raws, raw)
	}

	for _, raw := range raws {
		hash1, err := hashPassword(raw)
		if err != nil {
			t.Errorf(err.Error())
		}
		log.Println("hash1:", hash1)

		ok1, err := isConfigured(raw, hash1)
		if err != nil {
			t.Error(err)
		}
		log.Println("(ok1:", ok1, ")")

		hash2, err := hashPassword(raw)
		if err != nil {
			t.Errorf(err.Error())
		}
		log.Println("hash2:", hash2)

		ok2, err := isConfigured(raw, hash2)
		if err != nil {
			t.Error(err)
		}
		log.Println("(ok2:", ok2, ")")

		/* ng checker */
		ngHash, err := hashPassword("dummy")
		log.Println("ng hash:", ngHash)

		ng, err := isConfigured(raw, ngHash)
		log.Println("(false?:", ng, ")")
	}
}