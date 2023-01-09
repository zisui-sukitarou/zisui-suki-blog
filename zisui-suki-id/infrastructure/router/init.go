package router

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"zisui-suki-blog/adapter/controller"
	"zisui-suki-blog/infrastructure/db"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() (*echo.Echo, error) {
	e := echo.New()

	db, err := db.Init()
	if err != nil {
		return e, err
	}

	ctx := context.Background()

	user := controller.NewUserController(db)

	e.POST("/login", user.Login(&ctx))

	find := e.Group("/find")
	find.GET("/by/id", user.FindById(&ctx))
	find.GET("/by/name", user.FindByName(&ctx))

	findByToken := e.Group("/find/by/token")
	findByToken.Use(middleware.JWTWithConfig(newConfig()))
	findByToken.GET("", user.FindByToken(&ctx))

	log.Println("infra: router init")

	return e, nil
}

/* JWT config */
func getSecretKey() string {
	key := os.Getenv("JWT_SECRET_KEY")
	if key == "" {
		log.Panic("env: JWT_SECRET_KEY not specified")
	}
	return key
}

func newConfig() middleware.JWTConfig {
	return middleware.JWTConfig{
		SigningKey: []byte(getSecretKey()),
		ParseTokenFunc: func(tokenString string, c echo.Context) (interface{}, error) {
			keyFunc := func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(getSecretKey()), nil
			}

			token, err := jwt.Parse(tokenString, keyFunc)
			if err != nil {
				return nil, err
			}
			if !token.Valid {
				return nil, errors.New("invalid token")
			}
			return token, nil
		},
	}
}
