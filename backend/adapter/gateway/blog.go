package gateway

import (
	"context"
	"zisui-suki-blog/domain/model"
	"zisui-suki-blog/domain/repository"
	"zisui-suki-blog/ent"
	"zisui-suki-blog/ent/blog"
	"zisui-suki-blog/ent/tag"
	"zisui-suki-blog/ent/user"
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

func (g *BlogGateway) FindById(blogId model.BlogId) (bool, *repository.BlogData, error) {
	blog, err := g.Client.Debug().Blog.
		Query().
		Where(blog.IDEQ(string(blogId))).
		WithWriter().
		WithTags().
		First(*g.Context)
	if err != nil {
		return false, &repository.BlogData{}, err
	}

	/* create writer data */
	userData := repository.NewUserData(
		blog.Edges.Writer.ID,
		blog.Edges.Writer.Name,
		blog.Edges.Writer.Email,
		blog.Edges.Writer.Icon,
		blog.Edges.Writer.CreatedAt,
	)

	/* create tags data */
	var tags []*repository.TagData
	for _, tag := range blog.Edges.Tags {
		tags = append(tags, repository.NewTagData(
			tag.ID,
			tag.CreatedAt,
		))
	}

	return true, repository.NewBlogData(
		blog.ID,
		blog.Content,
		blog.Title,
		blog.Abstract,
		userData,
		tags,
		blog.CreatedAt,
		blog.UpdatedAt,
	), nil
}

func (g *BlogGateway) FindByUserId(userId model.UserId, begin uint, end uint) ([]*repository.BlogOverviewData, error) {

	// blog.content 以外を取得
	blogs, err := g.Client.Debug().Blog.Query().
		Select(blog.FieldID).
		Select(blog.FieldTitle).
		Select(blog.FieldAbstract).
		Select(blog.FieldEvaluation).
		Select(blog.FieldUserID).
		Select(blog.FieldCreatedAt).
		Select(blog.FieldUpdatedAt).
		Where(blog.HasWriterWith(user.IDEQ(string(userId)))).
		WithWriter().
		WithTags().
		Offset(int(begin)).
		Limit(int(end)).
		All(*g.Context)
	if err != nil {
		return []*repository.BlogOverviewData{}, err
	}

	var overviews []*repository.BlogOverviewData
	for _, blog := range blogs {
		/* create writer data */
		userData := repository.NewUserData(
			blog.Edges.Writer.ID,
			blog.Edges.Writer.Name,
			blog.Edges.Writer.Email,
			blog.Edges.Writer.Icon,
			blog.Edges.Writer.CreatedAt,
		)
		/* create tags data */
		var tags []*repository.TagData
		for _, tag := range blog.Edges.Tags {
			tags = append(tags, repository.NewTagData(
				tag.ID,
				tag.CreatedAt,
			))
		}
		/* create blog overviews */
		overviews = append(overviews, repository.NewBlogOverviewData(
			blog.ID,
			blog.Title,
			blog.Abstract,
			userData,
			tags,
			blog.CreatedAt,
			blog.UpdatedAt,
		))
	}
	return overviews, nil
}

func (g *BlogGateway) FindByTagName(tagName model.TagName, begin uint, end uint) ([]*repository.BlogOverviewData, error) {

	// blog.content 以外を取得
	blogs, err := g.Client.Debug().Blog.Query().
		Select(blog.FieldID).
		Select(blog.FieldTitle).
		Select(blog.FieldAbstract).
		Select(blog.FieldEvaluation).
		Select(blog.FieldUserID).
		Select(blog.FieldCreatedAt).
		Select(blog.FieldUpdatedAt).
		Where(blog.HasTagsWith(tag.IDEQ(string(tagName)))).
		WithWriter().
		WithTags().
		Offset(int(begin)).
		Limit(int(end)).
		All(*g.Context)
	if err != nil {
		return []*repository.BlogOverviewData{}, err
	}

	var overviews []*repository.BlogOverviewData
	for _, blog := range blogs {
		/* create writer data */
		userData := repository.NewUserData(
			blog.Edges.Writer.ID,
			blog.Edges.Writer.Name,
			blog.Edges.Writer.Email,
			blog.Edges.Writer.Icon,
			blog.Edges.Writer.CreatedAt,
		)
		/* create tags data */
		var tags []*repository.TagData
		for _, tag := range blog.Edges.Tags {
			tags = append(tags, repository.NewTagData(
				tag.ID,
				tag.CreatedAt,
			))
		}
		/* create blog overviews */
		overviews = append(overviews, repository.NewBlogOverviewData(
			blog.ID,
			blog.Title,
			blog.Abstract,
			userData,
			tags,
			blog.CreatedAt,
			blog.UpdatedAt,
		))
	}
	return overviews, nil
}

func (g *BlogGateway) FindByUserIdAndTagName(userId model.UserId, tagName model.TagName, begin uint, end uint) ([]*repository.BlogOverviewData, error) {

	// blog.content 以外を取得
	blogs, err := g.Client.Debug().Blog.Query().
		Select(blog.FieldID).
		Select(blog.FieldTitle).
		Select(blog.FieldAbstract).
		Select(blog.FieldEvaluation).
		Select(blog.FieldUserID).
		Select(blog.FieldCreatedAt).
		Select(blog.FieldUpdatedAt).
		Where(blog.HasWriterWith(user.IDEQ(string(userId)))).
		Where(blog.HasTagsWith(tag.IDEQ(string(tagName)))).
		WithWriter().
		WithTags().
		Offset(int(begin)).
		Limit(int(end)).
		All(*g.Context)
	if err != nil {
		return []*repository.BlogOverviewData{}, err
	}

	var overviews []*repository.BlogOverviewData
	for _, blog := range blogs {
		/* create writer data */
		userData := repository.NewUserData(
			blog.Edges.Writer.ID,
			blog.Edges.Writer.Name,
			blog.Edges.Writer.Email,
			blog.Edges.Writer.Icon,
			blog.Edges.Writer.CreatedAt,
		)
		/* create tags data */
		var tags []*repository.TagData
		for _, tag := range blog.Edges.Tags {
			tags = append(tags, repository.NewTagData(
				tag.ID,
				tag.CreatedAt,
			))
		}
		/* create blog overviews */
		overviews = append(overviews, repository.NewBlogOverviewData(
			blog.ID,
			blog.Title,
			blog.Abstract,
			userData,
			tags,
			blog.CreatedAt,
			blog.UpdatedAt,
		))
	}
	return overviews, nil
}

func (g *BlogGateway) Register(blog *model.Blog) error {
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

func (g *BlogGateway) RegisterTags(blogId model.BlogId, tagNames []model.TagName) error {
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

func (g *BlogGateway) Update(blog *model.Blog) (*repository.BlogData, error) {
	return &repository.BlogData{}, nil
}

func (g *BlogGateway) Delete(blog *model.Blog) error {
	return nil
}
