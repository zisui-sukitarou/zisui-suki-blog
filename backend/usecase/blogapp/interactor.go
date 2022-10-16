package blogapp

import (
	"errors"
	"time"
	"zisui-suki-blog/domain/model"
	"zisui-suki-blog/domain/repository"
	"zisui-suki-blog/domain/service"
	"zisui-suki-blog/usecase/apperr"
)

type BlogInteractor struct {
	OutputPort BlogOutputPort
	Repository repository.BlogRepository
}

/* constructor */
func NewBlogInteractor(
	outputPort BlogOutputPort,
	repo repository.BlogRepository,
) BlogInputPort {
	return &BlogInteractor{
		OutputPort: outputPort,
		Repository: repo,
	}
}

func (b *BlogInteractor) FindById(request *BlogFindByIdRequest) error {
	/* input data ->　model object */
	blogId, err := model.NewBlogId(request.BlogId)
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* find by id */
	exists, blog, err := b.Repository.FindById(blogId)
	if !exists {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(errors.New("blog of the id dosen't exist")))
	}
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* response */
	return b.OutputPort.RespondBlog(NewBlogResponse(blog))
}

func (b *BlogInteractor) FindByUserId(request *BlogFindByUserIdRequest) error {
	/* input data -> model object */
	userId, err := model.NewUserId(request.UserId)
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* find by user id */
	blogs, err := b.Repository.FindByUserId(userId, request.Begin, request.End)
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* response */
	return b.OutputPort.RespondBlogOverviews(NewBlogOverviewsResponse(blogs))
}

func (b *BlogInteractor) FindByTagName(request *BlogFindByTagRequest) error {
	/* input data -> model object */
	tagName, err := model.NewTagName(request.TagName)
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* find by tag_name */
	blogs, err := b.Repository.FindByTagName(tagName, request.Begin, request.End)
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* response */
	return b.OutputPort.RespondBlogOverviews(NewBlogOverviewsResponse(blogs))
}

func (b *BlogInteractor) FindByUserIdAndTagName(request *BlogFindByUserIdAndTagRequest) error {
	/* input data -> model object */
	userId, err := model.NewUserId(request.UserId)
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	tagName, err := model.NewTagName(request.TagName)
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* find by user_id & tag_name */
	blogs, err := b.Repository.FindByUserIdAndTagName(userId, tagName, request.Begin, request.End)
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* response */
	return b.OutputPort.RespondBlogOverviews(NewBlogOverviewsResponse(blogs))
}

func (b *BlogInteractor) Delete(request *BlogDeleteRequest) error {
	/* input data -> model object */
	blogId, err := model.NewBlogId(request.BlogId)
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* find by id */
	exists, blog, err := b.Repository.FindById(blogId)
	if !exists {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(errors.New("blog of the id dosen't exist")))
	}
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* delete */
	return b.Repository.Delete(blog)
}

func (b *BlogInteractor) Update(request *BlogUpdateRequest) error {
	/* input data -> madel object */
	blogId, err := model.NewBlogId(request.BlogId)
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	content, err := model.NewBlogContent(request.Content)
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	abstract, err := model.NewBlogAbstract(request.Abstract)
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	title, err := model.NewBlogTitle(request.Title)
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	evaluation, err := model.NewBlogEvaluation(request.Evaluation)
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* find */
	exists, blog, err := b.Repository.FindById(blogId)
	if !exists {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(errors.New("blog of the id dosen't exist")))
	}
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* update & save */
	blog.Update(content, title, abstract, evaluation)
	blog, err = b.Repository.Update(blog)
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* response */
	return b.OutputPort.RespondBlog(NewBlogResponse(blog))
}

func (b *BlogInteractor) Register(request *BlogRegisterRequest) error {
	/* input data -> model object */
	userId, err := model.NewUserId(request.UserId)
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	content, err := model.NewBlogContent(request.Content)
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	title, err := model.NewBlogTitle(request.Title)
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	abstract, err := model.NewBlogAbstract(request.Abstract)
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	evaluation, err := model.NewBlogEvaluation(request.Evaluation)
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	var tagNames []model.TagName
	for _, tag := range request.Tags {
		tagName, err := model.NewTagName(tag)
		if err != nil {
			return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
		}
		tagNames = append(tagNames, tagName)
	}

	/* gen blog_id */
	id, err := b.newService().GenULID()
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	blogId, err := model.NewBlogId(id.String())
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* create blog & save */
	blog := model.NewBlog(blogId, userId, content, title, abstract, evaluation, time.Now(), time.Now())
	err = b.Repository.Register(blog)
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* register tags info */
	err = b.Repository.RegisterTags(blogId, tagNames)
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* response */
	return b.OutputPort.RespondBlog(NewBlogResponse(blog))
}

/* create blog service */
func (b *BlogInteractor) newService() *service.Blog {
	return &service.Blog{}
}
