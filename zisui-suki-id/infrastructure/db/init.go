package db

import (
	"log"
	"os"
	"zisui-suki-blog/ent"

	"entgo.io/ent/dialect/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Init() (*ent.Client, error) {
	return open()
}

func getDbURL() string {
	url := os.Getenv("ID_DB_URL")
	if url == "" {
		log.Panic("env: ID_DB_URL not specified")
	}
	return url
}

func open() (*ent.Client, error) {
	drv, err := sql.Open("mysql", getDbURL())
	if err != nil {
		return nil, err
	}
	// Get the underlying sql.DB object of the driver.
	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	return ent.NewClient(ent.Driver(drv)), nil
}
