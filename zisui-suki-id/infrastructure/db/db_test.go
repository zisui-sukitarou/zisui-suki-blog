package db_test

import (
	"context"
	"log"
	"testing"
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

	ctx := context.Background()

	if err := db.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithForeignKeys(true),
	); err != nil {
		t.Errorf(err.Error())
	}
}