package db

import (
	"zisui-suki-blog/ent"

	"entgo.io/ent/dialect/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Init() (*ent.Client, error) {
	return open()
}

func open() (*ent.Client, error) {
	drv, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/id_db?charset=utf8&parseTime=true")
	if err != nil {
		return nil, err
	}
	// Get the underlying sql.DB object of the driver.
	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	return ent.NewClient(ent.Driver(drv)), nil
}
