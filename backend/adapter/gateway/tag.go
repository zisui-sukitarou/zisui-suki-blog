package gateway

import (
	"context"
	"zisui-suki-blog/domain/model"
	"zisui-suki-blog/domain/repository"
	"zisui-suki-blog/ent"
	"zisui-suki-blog/ent/tag"
)

type TagGateway struct {
	client  *ent.Client
	context *context.Context
}

func NewTagGateway(
	client *ent.Client,
	context *context.Context,
) repository.TagRepository {
	return &TagGateway{
		client:  client,
		context: context,
	}
}

func (g *TagGateway) Register(tag *model.Tag) error {
	_, err := g.client.Debug().Tag.Create().
		SetID(string(tag.TagName)).
		SetCreatedAt(tag.CreatedAt).
		Save(*g.context)
	return err
}

func (g *TagGateway) FindByTagName(tagName model.TagName) (bool, *model.Tag, error) {
	tag, err := g.client.Debug().Tag.Query().
		Where(tag.IDEQ(string(tagName))).
		First(*g.context)
	if err != nil {
		return false, &model.Tag{}, err
	}

	return true, model.NewTag(
		model.TagName(tag.ID),
		tag.CreatedAt,
	), nil
}
