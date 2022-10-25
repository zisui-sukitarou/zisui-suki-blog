package main

import (
	"log"
	"zisui-suki-blog/infrastructure/router"
)


func main() {
	r, err := router.Init()
	if err != nil {
		log.Println(err)
	}

	r.Start(":3001")
}