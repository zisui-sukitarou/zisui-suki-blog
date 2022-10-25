package router

import (
	"context"
	"errors"
	"fmt"
	"zisui-suki-blog/adapter/controller"
	"zisui-suki-blog/infrastructure/api"
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

	api, err := api.Init()
	if err != nil {
		return e, err
	}

	ctx := context.Background()
	blog := controller.NewBlogController(db, api.AuthApiDomain)

	find := e.Group("/find")
	find.GET("/by/id", blog.FindById(&ctx))
	find.GET("/by/tag", blog.FindByTagName(&ctx))
	find.GET("/by/user", blog.FindByUserId(&ctx))
	find.GET("/by/tag/user", blog.FindByUserIdAndTagName(&ctx))

	register := e.Group("/register")
	register.Use(middleware.JWTWithConfig(newConfig()))
	register.POST("", blog.Register(&ctx))

	update := e.Group("/update")
	update.Use(middleware.JWTWithConfig(newConfig()))
	update.POST("", blog.Update(&ctx))

	return e, nil
}

/* JWT config */
func newConfig() middleware.JWTConfig {
	return middleware.JWTConfig{
		SigningKey: []byte("SECRET_KEY"),
		ParseTokenFunc: func(tokenString string, c echo.Context) (interface{}, error) {
			keyFunc := func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte("SECRET_KEY"), nil
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
