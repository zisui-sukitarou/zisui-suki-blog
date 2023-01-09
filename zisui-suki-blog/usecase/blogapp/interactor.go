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
	BlogRepo   repository.BlogRepository
	UserRepo   repository.UserRepository
	TagRepo    repository.TagRepository
}

/* constructor */
func NewBlogInteractor(
	outputPort BlogOutputPort,
	blogRepo repository.BlogRepository,
	userRepo repository.UserRepository,
	tagRepo repository.TagRepository,
) BlogInputPort {
	return &BlogInteractor{
		OutputPort: outputPort,
		BlogRepo:   blogRepo,
		UserRepo:   userRepo,
		TagRepo:    tagRepo,
	}
}

func (b *BlogInteractor) FindById(request *BlogFindByIdRequest) error {
	/* input data ->　model object */
	blogId, err := model.NewBlogId(request.BlogId)
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* find blog by id */
	exists, blog, err := b.BlogRepo.FindById(blogId)
	if !exists {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(errors.New("blog of the id dosen't exist")))
	}
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* find user by user_id */
	exists, user, err := b.UserRepo.FindById(blog.Blog.UserId)
	if !exists {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(errors.New("user of the id dosen't exist")))
	}
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* response */
	return b.OutputPort.FindById(NewBlogResponse(blog, user))
}

func (b *BlogInteractor) FindByUserId(request *BlogFindByUserIdRequest) error {
	/* input data -> model object */
	userId, err := model.NewUserId(request.UserId)
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* find by user id */
	blogs, err := b.BlogRepo.FindByUserId(userId, request.Begin, request.End)
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	var response []*BlogOverviewResponse
	for _, blog := range blogs {
		/* find user by user_id */
		exists, user, err := b.UserRepo.FindById(blog.Blog.UserId)
		if !exists {
			return b.OutputPort.RespondErorr(apperr.NewErrorResponse(errors.New("user of the id dosen't exist")))
		}
		if err != nil {
			return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
		}

		response = append(response, NewBlogOverviewResponse(blog, user))
	}

	/* response */
	return b.OutputPort.FindByUserId(response)
}

func (b *BlogInteractor) FindByUserName(request *BlogFindByUserNameRequest) error {
	/* input data -> model object */
	userName, err := model.NewUserName(request.UserName)
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* userName -> userId */
	exists, user, err := b.UserRepo.FindByUserName(userName)
	if !exists || err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* find by user id */
	blogs, err := b.BlogRepo.FindByUserId(user.UserId, request.Begin, request.End)
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	var response []*BlogOverviewResponse
	for _, blog := range blogs {
		/* find user by user_id */
		exists, user, err := b.UserRepo.FindById(blog.Blog.UserId)
		if !exists {
			return b.OutputPort.RespondErorr(apperr.NewErrorResponse(errors.New("user of the id dosen't exist")))
		}
		if err != nil {
			return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
		}

		response = append(response, NewBlogOverviewResponse(blog, user))
	}

	/* response */
	return b.OutputPort.FindByUserName(response)
}

func (b *BlogInteractor) FindByTagName(request *BlogFindByTagRequest) error {
	/* input data -> model object */
	tagName, err := model.NewTagName(request.TagName)
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* find by tag_name */
	blogs, err := b.BlogRepo.FindByTagName(tagName, request.Begin, request.End)
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	var response []*BlogOverviewResponse
	for _, blog := range blogs {
		/* find user by user_id */
		exists, user, err := b.UserRepo.FindById(blog.Blog.UserId)
		if !exists {
			return b.OutputPort.RespondErorr(apperr.NewErrorResponse(errors.New("user of the id dosen't exist")))
		}
		if err != nil {
			return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
		}

		response = append(response, NewBlogOverviewResponse(blog, user))
	}

	/* response */
	return b.OutputPort.FindByTagName(response)
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
	blogs, err := b.BlogRepo.FindByUserIdAndTagName(userId, tagName, request.Begin, request.End)
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	var response []*BlogOverviewResponse
	for _, blog := range blogs {
		/* find user by user_id */
		exists, user, err := b.UserRepo.FindById(blog.Blog.UserId)
		if !exists {
			return b.OutputPort.RespondErorr(apperr.NewErrorResponse(errors.New("user of the id dosen't exist")))
		}
		if err != nil {
			return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
		}

		response = append(response, NewBlogOverviewResponse(blog, user))
	}

	/* response */
	return b.OutputPort.FindByUserIdAndTagName(response)
}

func (b *BlogInteractor) FindByUserNameAndTagName(request *BlogFindByUserNameAndTagRequest) error {
	/* input data -> model object */
	userName, err := model.NewUserName(request.UserName)
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	tagName, err := model.NewTagName(request.TagName)
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* userName -> userId */
	exists, user, err := b.UserRepo.FindByUserName(userName)
	if !exists || err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* find by user id */
	blogs, err := b.BlogRepo.FindByUserIdAndTagName(user.UserId, tagName, request.Begin, request.End)
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	var response []*BlogOverviewResponse
	for _, blog := range blogs {
		/* find user by user_id */
		exists, user, err := b.UserRepo.FindById(blog.Blog.UserId)
		if !exists {
			return b.OutputPort.RespondErorr(apperr.NewErrorResponse(errors.New("user of the id dosen't exist")))
		}
		if err != nil {
			return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
		}

		response = append(response, NewBlogOverviewResponse(blog, user))
	}

	/* response */
	return b.OutputPort.FindByUserNameAndTagName(response)
}

func (b *BlogInteractor) Delete(request *BlogDeleteRequest) error {
	/* input data -> model object */
	blogId, err := model.NewBlogId(request.BlogId)
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* find by id */
	exists, blog, err := b.BlogRepo.FindById(blogId)
	if err != nil {
		b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}
	if !exists {
		b.OutputPort.RespondErorr(apperr.NewErrorResponse(errors.New("blog of the id dosen't exist")))
	}

	/* delete */
	return b.BlogRepo.Delete(blog.Blog)
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

	var tagNames []model.TagName
	for _, tag := range request.Tags {
		tagName, err := model.NewTagName(tag)
		if err != nil {
			return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
		}
		tagNames = append(tagNames, tagName)

		// タグが未登録なら登録
		tag := model.NewTag(
			tagName,
			model.TagIcon(""),
			model.TagStatus(0),
			time.Now(),
			time.Now(),
		)
		b.TagRepo.Register(tag)
	}

	/* find blog by id */
	exists, blog, err := b.BlogRepo.FindById(blogId)
	if !exists {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(errors.New("blog of the id dosen't exist")))
	}
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* update blog & save */
	blog.Blog.Update(content, title, abstract, evaluation)
	err = b.BlogRepo.Update(blog.Blog)
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* update tags */
	err = b.BlogRepo.UpdateTags(blogId, tagNames)
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* find user by user_id */
	exists, user, err := b.UserRepo.FindById(blog.Blog.UserId)
	if !exists {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(errors.New("user of the id dosen't exist")))
	}
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* response */
	return b.OutputPort.Update(NewBlogResponse(blog, user))
}

func (b *BlogInteractor) Register(request *BlogRegisterRequest) error {
	/* input data -> model object */
	blogId, err := model.NewBlogId(request.BlogId)
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

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

	status, err := model.NewBlogStatus(request.Status)
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

		// タグが未登録なら登録
		tag := model.NewTag(
			tagName,
			model.TagIcon(""),
			model.TagStatus(0),
			time.Now(),
			time.Now(),
		)
		b.TagRepo.Register(tag)
	}

	/* create blog & save */
	blog := model.NewBlog(blogId, userId, content, title, abstract, evaluation, status, time.Now(), time.Now())
	err = b.BlogRepo.Register(blog)
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* register tags info */
	err = b.BlogRepo.RegisterTags(blogId, tagNames)
	if err != nil {
		return b.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* response */
	return b.OutputPort.Register()
}

/* create blog service */
func (b *BlogInteractor) service() *service.Blog {
	return &service.Blog{}
}

/* model.TagName[] contains model.TagName ?*/
func (b *BlogInteractor) contain(tags []*model.Tag, target model.TagName) bool {
	for _, tag := range tags {
		if tag.TagName == target {
			return true
		}
	}
	return false
}
