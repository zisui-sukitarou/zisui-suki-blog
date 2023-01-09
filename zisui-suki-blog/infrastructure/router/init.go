package router

import (
	"context"
	"errors"
	"fmt"
	"os"
	"log"
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

	/* public blog */
	blog := controller.NewBlogController(db, api.IdApiBaseURL)

	findBlog := e.Group("/find/blog")
	findBlog.GET("/by/id", blog.FindById(&ctx))
	findBlog.GET("/by/tag", blog.FindByTagName(&ctx))
	findBlog.GET("/by/userName", blog.FindByUserName(&ctx))
	findBlog.GET("/by/tag/userName", blog.FindByUserNameAndTagName(&ctx))

	registerBlog := e.Group("/register/blog")
	registerBlog.Use(middleware.JWTWithConfig(newConfig()))
	registerBlog.POST("", blog.Register(&ctx))

	updateBlog := e.Group("/update/blog")
	updateBlog.Use(middleware.JWTWithConfig(newConfig()))
	updateBlog.POST("", blog.Update(&ctx))

	/* private draft */
	draft := controller.NewDraftController(db, api.IdApiBaseURL)

	findDraft := e.Group("/find/draft")
	findDraft.Use(middleware.JWTWithConfig(newConfig()))
	findDraft.GET("/by/id", draft.FindById(&ctx))
	findDraft.GET("/by/userId", draft.FindByUserId(&ctx))

	newDraft := e.Group("/new/draft")
	newDraft.Use(middleware.JWTWithConfig(newConfig()))
	newDraft.POST("", draft.New(&ctx))

	registerDraft := e.Group("/register/draft")
	registerDraft.Use(middleware.JWTWithConfig(newConfig()))
	registerDraft.POST("", draft.Register(&ctx))

	updateDraft := e.Group("/update/draft")
	updateDraft.Use(middleware.JWTWithConfig(newConfig()))
	updateDraft.POST("", draft.Update(&ctx))

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
