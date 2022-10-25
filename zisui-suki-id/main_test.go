package main_test

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"

	"encoding/json"
)

type userFindByName struct {
	Name string `json:"name"`
}

type userData struct {
	UserId      string    `json:"user_id"`
	Email       string    `json:"email"`
	Name        string    `json:"name"`
	DisplayName string    `json:"display_name"`
	Icon        string    `json:"icon"`
	Status      uint      `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func TestRequest(t *testing.T) {
	/* json */
	zisui := userFindByName{Name: "zisui-sukitarou"}
	zisuiJson, _ := json.Marshal(zisui)

	/* construct http request */
	client := http.Client{}

	req, err := http.NewRequest("GET", "http://localhost:3001/find/by/name", bytes.NewBuffer(zisuiJson))
	if err != nil {
		t.Error(err)
	}

	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMDFHRzRFQlZHQ0Y2WkM1TkU5MUtGWjZQSDUifQ.KhVnae0iF7uLZWPxfAcC3aIzlAmkiO4jMMTrFb3OgcQ")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("accept", "application/json")

	/* do request */
	response, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}

	user := userData{}
	if err := json.Unmarshal(body, &user); err != nil {
		t.Error(err)
	}
	log.Println("user:", user)
}
