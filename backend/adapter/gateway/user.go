package gateway

import (
	"context"
	"zisui-suki-blog/domain/model"
	"zisui-suki-blog/domain/repository"
	"zisui-suki-blog/ent"
	"zisui-suki-blog/ent/user"
)

type UserGateway struct {
	Client  *ent.Client
	Context *context.Context
}

func NewUserGateway(
	client *ent.Client,
	ctx *context.Context,
) repository.UserRepository {
	return &UserGateway{
		Client:  client,
		Context: ctx,
	}
}

func (g *UserGateway) FindById(userId model.UserId) (bool, *repository.UserData, error) {
	user, err := g.Client.Debug().User.
		Query().
		Where(user.IDEQ(string(userId))).
		First(*g.Context)
	if err != nil {
		return false, &repository.UserData{}, err
	}

	return true, repository.NewUserData(
		user.ID,
		user.Name,
		user.Email,
		user.Icon,
		user.CreatedAt,
	), nil
}

func (g *UserGateway) FindByEmail(email model.UserEmail) (bool, *repository.UserData, error) {
	user, err := g.Client.Debug().User.
		Query().
		Where(user.EmailEQ(string(email))).
		First(*g.Context)
	if err != nil {
		return false, &repository.UserData{}, err
	}

	return true, repository.NewUserData(
		user.ID,
		user.Name,
		user.Email,
		user.Icon,
		user.CreatedAt,
	), nil
}

func (g *UserGateway) Register(user *model.User) error {
	_, err := g.Client.Debug().User.Create().
		SetID(string(user.UserId)).
		SetName(string(user.Name)).
		SetPassword(string(user.HashedPassword)).
		SetEmail(string(user.Email)).
		SetIcon(string(user.Icon)).
		SetCreatedAt(user.CreatedAt).
		SetUpdatedAt(user.UpdatedAt).
		Save(*g.Context)
	return err
}

func (g *UserGateway) Update(nInfo *model.User) (*repository.UserData, error) {
	user, err := g.Client.Debug().User.
		UpdateOneID(string(nInfo.UserId)).
		SetName(string(nInfo.Name)).
		SetEmail(string(nInfo.Email)).
		SetIcon(string(nInfo.Icon)).
		Save(*g.Context)
	if err != nil {
		return &repository.UserData{}, err
	}

	return repository.NewUserData(
		user.ID,
		user.Name,
		user.Email,
		user.Icon,
		user.CreatedAt,
	), nil
}

func (g *UserGateway) Delete(user *model.User) error {
	return g.Client.Debug().User.
		DeleteOneID(string(user.UserId)).
		Exec(*g.Context)
}
