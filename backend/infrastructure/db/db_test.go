package db_test

import (
	"log"
	"testing"
	"time"
	"zisui-suki-blog/ent/migrate"
	"zisui-suki-blog/infrastructure/db"
)

func TestInitDB(t *testing.T) {
	db, err := db.Init()
	if err != nil {
		t.Errorf(err.Error())
	}
	log.Print(db)
}

func TestMigration(t *testing.T) {
	db, err := db.Init()
	if err != nil {
		t.Errorf(err.Error())
	}

	client := db.Client
	ctx := db.Ctx

	if err := client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithForeignKeys(true),
	); err != nil {
		t.Errorf(err.Error())
	}
}

func TestInsert1(t *testing.T) {
	db, err := db.Init()
	if err != nil {
		t.Errorf(err.Error())
	}

	client := db.Client
	ctx := db.Ctx

	user, err := client.User.Create().
		SetID("zisui-sukitarou").
		SetName("Tarou2").
		SetEmail("yyy.gmail.com").
		SetIcon("fuga.jpg").
		SetPassword("fuga").
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		t.Errorf(err.Error())
	}
	log.Print("inserted:", user)
}