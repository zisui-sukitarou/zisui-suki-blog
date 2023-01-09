package main

import (
	"log"
	"zisui-suki-blog/infrastructure/router"
	"os"
)

func getPort() (string) {
	port := os.Getenv("BLOG_PORT")
	if port == "" {
		log.Panic("env: BLOG_PORT not specified")
	}
	return port
}

func main() {
	r, err := router.Init()
	if err != nil {
		log.Panic(err)
	}

	r.Start(getPort())
}