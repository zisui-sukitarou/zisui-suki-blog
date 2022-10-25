package gateway

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
	"zisui-suki-blog/domain/model"
	"zisui-suki-blog/domain/repository"

	"github.com/labstack/echo/v4"
)

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
	UserId      string    `json:"user_id"`
	Email       string    `json:"email"`
	Name        string    `json:"name"`
	DisplayName string    `json:"display_name"`
	Icon        string    `json:"icon"`
	Status      uint      `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

/* find by id */
type userFindById struct {
	UserId string `json:"user_id"`
}

func (g *UserGateway) FindById(userId model.UserId) (bool, *model.User, error) {
	/* model -> req json */
	reqData := userFindById{
		UserId: string(userId),
	}
	reqJson, err := json.Marshal(reqData)
	if err != nil {
		return false, &model.User{}, err
	}

	/* extract token from echo context */
	token := g.ctx.Request().Header.Get("Authorization")

	/* do http request -> get response */
	client := http.Client{}
	url := g.domain + "find/by/id"
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(reqJson))
	if err != nil {
		return false, &model.User{}, err
	}

	req.Header.Add("Authorization", token)
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
