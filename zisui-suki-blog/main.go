package main

import (
	"log"
	"zisui-suki-blog/infrastructure/router"
)


func main() {
	r, err := router.Init()
	if err != nil {
		log.Panic(err)
	}

	r.Start(":3002")
}