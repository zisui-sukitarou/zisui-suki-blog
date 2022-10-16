package db

import (
	"context"
	"time"
	"zisui-suki-blog/ent"

	"entgo.io/ent/dialect/sql"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	Ctx    context.Context
	Client *ent.Client
}

func Init() (*DB, error) {
	ctx := context.Background()
	client, err := open()
	if err != nil {
		return &DB{}, err
	}

	return &DB{
		Ctx:    ctx,
		Client: client,
	}, nil
}

func open() (*ent.Client, error) {
	drv, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/zisuidb?charset=utf8&parseTime=true")
	if err != nil {
		return nil, err
	}
	// Get the underlying sql.DB object of the driver.
	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	return ent.NewClient(ent.Driver(drv)), nil
}
