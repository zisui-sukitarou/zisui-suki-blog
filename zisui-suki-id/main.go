package main

import (
	"log"
	"zisui-suki-blog/infrastructure/router"
	"os"
)

func getPort() (string) {
	port := os.Getenv("ID_PORT")
	if port == "" {
		log.Panic("env: ID_PORT not specified")
	}
	return port
}

func main() {
	r, err := router.Init()
	if err != nil {
		log.Println(err)
	}

	r.Start(getPort())
}