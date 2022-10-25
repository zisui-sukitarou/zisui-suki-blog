package gateway

import (
	"context"
	"zisui-suki-blog/domain/model"
	"zisui-suki-blog/domain/repository"
	"zisui-suki-blog/ent"
	"zisui-suki-blog/ent/blog"
	"zisui-suki-blog/ent/tag"
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
		WithTags().
		First(*g.Context)
	if err != nil {
		return false, &repository.BlogData{}, err
	}

	/* create blog model object */
	b := model.NewBlog(
		model.BlogId(blog.ID),
		model.UserId(blog.UserID),
		model.BlogContent(blog.Content),
		model.BlogTitle(blog.Title),
		model.BlogAbstract(blog.Abstract),
		model.BlogEvaluation(blog.Evaluation),
		model.BlogStatus(blog.Status),
		blog.CreatedAt,
		blog.UpdatedAt,
	)

	/* create tags model object */
	var tags []*model.Tag
	for _, tag := range blog.Edges.Tags {
		tags = append(tags, model.NewTag(
			model.TagName(tag.ID),
				model.TagIcon(tag.Icon),
				model.TagStatus(tag.Status),
				tag.CreatedAt,
				tag.UpdatedAt,
		))
	}

	return true, repository.NewBlogData(b, tags), nil
}

func (g *BlogGateway) FindByUserId(userId model.UserId, begin uint, end uint) ([]*repository.BlogData, error) {

	// blog.content 以外を取得
	blogs, err := g.Client.Debug().Blog.Query().
		Select(blog.FieldID).
		Select(blog.FieldUserID).
		Select(blog.FieldTitle).
		Select(blog.FieldAbstract).
		Select(blog.FieldEvaluation).
		Select(blog.FieldStatus).
		Select(blog.FieldCreatedAt).
		Select(blog.FieldUpdatedAt).
		Where(blog.UserIDEQ(string(userId))).
		WithTags().
		Offset(int(begin)).
		Limit(int(end)).
		All(*g.Context)
	if err != nil {
		return []*repository.BlogData{}, err
	}

	var overviews []*repository.BlogData
	for _, blog := range blogs {
		/* create blog model object */
		b := model.NewBlog(
			model.BlogId(blog.ID),
			model.UserId(blog.UserID),
			model.BlogContent(blog.Content),
			model.BlogTitle(blog.Title),
			model.BlogAbstract(blog.Abstract),
			model.BlogEvaluation(blog.Evaluation),
			model.BlogStatus(blog.Status),
			blog.CreatedAt,
			blog.UpdatedAt,
		)

		/* create tags model object */
		var tags []*model.Tag
		for _, tag := range blog.Edges.Tags {
			tags = append(tags, model.NewTag(
				model.TagName(tag.ID),
				model.TagIcon(tag.Icon),
				model.TagStatus(tag.Status),
				tag.CreatedAt,
				tag.UpdatedAt,
			))
		}

		/* create blog overviews */
		overviews = append(overviews, repository.NewBlogData(b, tags))
	}
	return overviews, nil
}

func (g *BlogGateway) FindByTagName(tagName model.TagName, begin uint, end uint) ([]*repository.BlogData, error) {

	// blog.content 以外を取得
	blogs, err := g.Client.Debug().Blog.Query().
		Select(blog.FieldID).
		Select(blog.FieldTitle).
		Select(blog.FieldAbstract).
		Select(blog.FieldEvaluation).
		Select(blog.FieldStatus).
		Select(blog.FieldUserID).
		Select(blog.FieldCreatedAt).
		Select(blog.FieldUpdatedAt).
		Where(blog.HasTagsWith(tag.IDEQ(string(tagName)))).
		WithTags().
		Offset(int(begin)).
		Limit(int(end)).
		All(*g.Context)
	if err != nil {
		return []*repository.BlogData{}, err
	}

	var overviews []*repository.BlogData
	for _, blog := range blogs {
		/* create blog model object */
		b := model.NewBlog(
			model.BlogId(blog.ID),
			model.UserId(blog.UserID),
			model.BlogContent(blog.Content),
			model.BlogTitle(blog.Title),
			model.BlogAbstract(blog.Abstract),
			model.BlogEvaluation(blog.Evaluation),
			model.BlogStatus(blog.Status),
			blog.CreatedAt,
			blog.UpdatedAt,
		)

		/* create tags model object */
		var tags []*model.Tag
		for _, tag := range blog.Edges.Tags {
			tags = append(tags, model.NewTag(
				model.TagName(tag.ID),
				model.TagIcon(tag.Icon),
				model.TagStatus(tag.Status),
				tag.CreatedAt,
				tag.UpdatedAt,
			))
		}

		/* create blog overviews */
		overviews = append(overviews, repository.NewBlogData(b, tags))
	}
	return overviews, nil
}

func (g *BlogGateway) FindByUserIdAndTagName(userId model.UserId, tagName model.TagName, begin uint, end uint) ([]*repository.BlogData, error) {

	// blog.content 以外を取得
	blogs, err := g.Client.Debug().Blog.Query().
		Select(blog.FieldID).
		Select(blog.FieldTitle).
		Select(blog.FieldAbstract).
		Select(blog.FieldEvaluation).
		Select(blog.FieldStatus).
		Select(blog.FieldUserID).
		Select(blog.FieldCreatedAt).
		Select(blog.FieldUpdatedAt).
		Where(blog.UserIDEQ(string(userId))).
		Where(blog.HasTagsWith(tag.IDEQ(string(tagName)))).
		WithTags().
		Offset(int(begin)).
		Limit(int(end)).
		All(*g.Context)
	if err != nil {
		return []*repository.BlogData{}, err
	}

	var overviews []*repository.BlogData
	for _, blog := range blogs {
		/* create blog model object */
		b := model.NewBlog(
			model.BlogId(blog.ID),
			model.UserId(blog.UserID),
			model.BlogContent(blog.Content),
			model.BlogTitle(blog.Title),
			model.BlogAbstract(blog.Abstract),
			model.BlogEvaluation(blog.Evaluation),
			model.BlogStatus(blog.Status),
			blog.CreatedAt,
			blog.UpdatedAt,
		)

		/* create tags model object */
		var tags []*model.Tag
		for _, tag := range blog.Edges.Tags {
			tags = append(tags, model.NewTag(
				model.TagName(tag.ID),
				model.TagIcon(tag.Icon),
				model.TagStatus(tag.Status),
				tag.CreatedAt,
				tag.UpdatedAt,
			))
		}

		/* create blog overviews */
		overviews = append(overviews, repository.NewBlogData(b, tags))
	}
	return overviews, nil
}

func (g *BlogGateway) Register(blog *model.Blog) error {
	_, err := g.Client.Debug().Blog.Create().
		SetID(string(blog.BlogId)).
		SetUserID(string(blog.UserId)).
		SetTitle(string(blog.Title)).
		SetAbstract(string(blog.Abstract)).
		SetContent(string(blog.Content)).
		SetEvaluation(uint(blog.Evaluation)).
		SetStatus(uint(blog.Status)).
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
