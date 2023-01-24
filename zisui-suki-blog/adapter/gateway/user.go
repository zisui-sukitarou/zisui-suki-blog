package gateway

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"zisui-suki-blog/domain/model"
	"zisui-suki-blog/domain/repository"

	"github.com/labstack/echo/v4"
)

// TODO: キャッシュ使うように
type UserGateway struct {
	domain string
	ctx    echo.Context
}

func NewUserGateway(
	domain string,
	ctx echo.Context,
) repository.UserRepository {
	return &UserGateway{
		domain: domain,
		ctx:    ctx,
	}
}

type UserData struct {
	UserId      string    `json:"userId"`
	Email       string    `json:"email"`
	Name        string    `json:"name"`
	DisplayName string    `json:"displayName"`
	Icon        string    `json:"icon"`
	Status      uint      `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type UserIdData struct {
	UserId string `json:"userId"`
}

/* find by id */
func (g *UserGateway) FindById(userId model.UserId) (bool, *model.User, error) {

	/* do http request -> get response */
	client := http.Client{}
	baseUrl := g.domain + "find/by/id"
	query := "?userId=" + string(userId)
	req, err := http.NewRequest("GET", baseUrl+query, nil)
	if err != nil {
		return false, &model.User{}, err
	}

	// req.Header.Add("Authorization", token)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("accept", "application/json")

	response, err := client.Do(req)
	if err != nil {
		return false, &model.User{}, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false, &model.User{}, err
	}

	/* binding response into user_data */
	var user UserData
	if err := json.Unmarshal(body, &user); err != nil {
		return false, &model.User{}, err
	}

	return true, model.NewUser(
		model.UserId(user.UserId),
		model.UserName(user.Name),
		model.UserDisplayName(user.DisplayName),
		model.UserEmail(user.Email),
		model.UserIcon(user.Icon),
		model.UserStatus(user.Status),
		user.CreatedAt,
		user.UpdatedAt,
	), nil
}

/* find by id */
func (g *UserGateway) FindByUserName(userName model.UserName) (bool, *model.User, error) {

	/* do http request -> get response */
	client := http.Client{}
	baseUrl := g.domain + "find/by/name"
	query := "?name=" + string(userName)
	req, err := http.NewRequest("GET", baseUrl+query, nil)
	if err != nil {
		return false, &model.User{}, err
	}

	// req.Header.Add("Authorization", token)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("accept", "application/json")

	response, err := client.Do(req)
	if err != nil {
		return false, &model.User{}, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false, &model.User{}, err
	}

	/* binding response into user_data */
	var user UserData
	if err := json.Unmarshal(body, &user); err != nil {
		return false, &model.User{}, err
	}

	return true, model.NewUser(
		model.UserId(user.UserId),
		model.UserName(user.Name),
		model.UserDisplayName(user.DisplayName),
		model.UserEmail(user.Email),
		model.UserIcon(user.Icon),
		model.UserStatus(user.Status),
		user.CreatedAt,
		user.UpdatedAt,
	), nil
}

/* name -> id */
func (g *UserGateway) NameToId(userName model.UserName) (bool, model.UserId, error) {

	/* do http request -> get response */
	client := http.Client{}
	baseUrl := g.domain + "name/to/id"
	query := "?name=" + string(userName)
	log.Print(baseUrl+query)
	req, err := http.NewRequest("GET", baseUrl+query, nil)
	if err != nil {
		return false, model.UserId(""), err
	}

	// req.Header.Add("Authorization", token)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("accept", "application/json")

	response, err := client.Do(req)
	if err != nil {
		return false, model.UserId(""), err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false, model.UserId(""), err
	}

	/* binding response into user_data */
	var userId UserIdData
	if err := json.Unmarshal(body, &userId); err != nil {
		return false, model.UserId(""), err
	}

	return true, model.UserId(userId.UserId), nil
}
