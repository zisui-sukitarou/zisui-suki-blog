package gateway

import (
	"context"
	"zisui-suki-blog/domain/model"
	"zisui-suki-blog/domain/repository"
	"zisui-suki-blog/ent"
	"zisui-suki-blog/ent/blog"
)

type BlogGateway struct {
	Client  *ent.Client
	Context *context.Context
}

func NewBlogGateway(
	client *ent.Client,
	ctx *context.Context,
) repository.BlogRepository {
	return &BlogGateway{
		Client:  client,
		Context: ctx,
	}
}

func (g *BlogGateway) FindById(blogId model.BlogId) (bool, *model.Blog, error) {
	blog, err := g.Client.Debug().Blog.
		Query().
		Where(blog.IDEQ(string(blogId))).
		WithWriter().
		First(*g.Context)
	if err != nil {
		return false, &model.Blog{}, err
	}

	return true, model.NewBlog(
		model.BlogId(blog.ID),
		model.UserId(blog.UserID),
		model.BlogContent(blog.Content),
		model.BlogTitle(blog.Title),
		model.BlogAbstract(blog.Abstract),
		model.BlogEvaluation(blog.Evaluation),
		blog.CreatedAt,
		blog.CreatedAt,
	), nil
}

func (g *BlogGateway) FindByUserId(userId model.UserId) ([]*model.Blog, error) {

	return []*model.Blog{}, nil
}

func (g *BlogGateway) FindByTagName(tagName model.TagName) ([]*model.Blog, error) {

	return []*model.Blog{}, nil
}

func (g *BlogGateway) Register(blog *model.Blog) (error) {
	_, err := g.Client.Debug().Blog.Create().
		SetID(string(blog.BlogId)).
		SetWriterID(string(blog.UserId)).
		SetTitle(string(blog.Title)).
		SetAbstract(string(blog.Abstract)).
		SetContent(string(blog.Content)).
		SetEvaluation(uint(blog.Evaluation)).
		SetCreatedAt(blog.CreatedAt).
		SetUpdatedAt(blog.UpdatedAt).
		Save(*g.Context)
	return err
}

func (g *BlogGateway) RegisterTags(blogId model.BlogId, tagNames []model.TagName) (error) {
	/* tagNames ([]model.TagName) -> tagNameStrings ([]string) */
	var tagNameStrings []string
	for _, v := range tagNames {
		tagNameStrings = append(tagNameStrings, string(v))
	}

	/* save */
	_, err := g.Client.Debug().Blog.Update().
		Where(blog.IDEQ(string(blogId))).
		AddTagIDs(tagNameStrings...).
		Save(*g.Context)

	return err
}

func (g *BlogGateway) Update(blog *model.Blog) (*model.Blog, error) {
	return &model.Blog{}, nil
}

func (g *BlogGateway) Delete(blog *model.Blog) (error) {
	return nil
}