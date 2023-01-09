package api

import (
	"os"
	"log"
)

type API struct {
	IdApiBaseURL string
}

func getIdApiBaseURL() string {
	base := os.Getenv("ID_API_BASE_URL")
	if base == "" {
		log.Panic("env: ID_API_BASE_URL not specified")
	}
	return base
}

func Init() (*API, error) {
	return &API{
		IdApiBaseURL: getIdApiBaseURL(),
	}, nil
}
