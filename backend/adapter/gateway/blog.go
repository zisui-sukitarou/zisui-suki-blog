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

func (g *BlogGateway) FindByUserId(userId model.UserId, begin uint, end uint) ([]*model.Blog, error) {

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
		Offset(int(begin)).
		Limit(int(end)).
		All(*g.Context)
	if err != nil {
		return []*model.Blog{}, err
	}

	// content は不要なのでから文字列で代替
	var res []*model.Blog
	for _, blog := range blogs {
		res = append(res, model.NewBlog(
			model.BlogId(blog.ID),
			model.UserId(blog.UserID),
			model.BlogContent(""),
			model.BlogTitle(blog.Title),
			model.BlogAbstract(blog.Abstract),
			model.BlogEvaluation(blog.Evaluation),
			blog.CreatedAt,
			blog.UpdatedAt,
		))
	}
	return res, nil
}

func (g *BlogGateway) FindByTagName(tagName model.TagName, begin uint, end uint) ([]*model.Blog, error) {

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
		Offset(int(begin)).
		Limit(int(end)).
		All(*g.Context)
	if err != nil {
		return []*model.Blog{}, err
	}

	// content は不要なので空文字列で代替
	var res []*model.Blog
	for _, blog := range blogs {
		res = append(res, model.NewBlog(
			model.BlogId(blog.ID),
			model.UserId(blog.UserID),
			model.BlogContent(""),
			model.BlogTitle(blog.Title),
			model.BlogAbstract(blog.Abstract),
			model.BlogEvaluation(blog.Evaluation),
			blog.CreatedAt,
			blog.UpdatedAt,
		))
	}
	return res, nil
}

func (g *BlogGateway) FindByUserIdAndTagName(userId model.UserId, tagName model.TagName, begin uint, end uint) ([]*model.Blog, error) {

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
		Offset(int(begin)).
		Limit(int(end)).
		All(*g.Context)
	if err != nil {
		return []*model.Blog{}, err
	}

	// content は不要なので空文字列で代替
	var res []*model.Blog
	for _, blog := range blogs {
		res = append(res, model.NewBlog(
			model.BlogId(blog.ID),
			model.UserId(blog.UserID),
			model.BlogContent(""),
			model.BlogTitle(blog.Title),
			model.BlogAbstract(blog.Abstract),
			model.BlogEvaluation(blog.Evaluation),
			blog.CreatedAt,
			blog.UpdatedAt,
		))
	}
	return res, nil
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

func (g *BlogGateway) Update(blog *model.Blog) (*model.Blog, error) {
	return &model.Blog{}, nil
}

func (g *BlogGateway) Delete(blog *model.Blog) error {
	return nil
}
