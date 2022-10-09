package blogapp

import (
	"cook-blog/domain/model"
	"cook-blog/domain/repository"
	"cook-blog/domain/service"
	"errors"
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

func (b *BlogInteractor) FindById(request BlogFindByIdRequest) {
	/* input data ->　model object */
	blogId, err := model.NewBlogId(request.BlogId)
	if err != nil {
		b.OutputPort.RespondErorr(err)
	}

	/* find */
	exists, blog, err := b.Repository.FindById(blogId)
	if !exists {
		b.OutputPort.RespondErorr(errors.New("blog of the id dosen't exist"))
	}
	if err != nil {
		b.OutputPort.RespondErorr(err)
	}

	/* response */
	response := NewBlogResponse(blog)
	b.OutputPort.RespondBlog(response)
}

func (b *BlogInteractor) FindByUserId(request BlogFindByUserIdRequest) {
	/* input data -> model object */
	userId, err := model.NewUserId(request.UserId)
	if err != nil {
		b.OutputPort.RespondErorr(err)
	}

	/* find */
	blogs, err := b.Repository.FindByUserId(userId)
	if err != nil {
		b.OutputPort.RespondErorr(err)
	}

	/* response */
	response := NewBlogsResponse(blogs)
	b.OutputPort.RespondBlogs(response)
}

func (b *BlogInteractor) FindByTagId(request BlogFindByTagIdRequest) {

}

func (b *BlogInteractor) Delete(request BlogDeleteRequest) {
	/* input data -> model object */
	blogId, err := model.NewBlogId(request.BlogId)
	if err != nil {
		b.OutputPort.RespondErorr(err)
	}

	/* find */
	exists, blog, err := b.Repository.FindById(blogId)
	if !exists {
		b.OutputPort.RespondErorr(errors.New("blog of the id dosen't exist"))
	}
	if err != nil {
		b.OutputPort.RespondErorr(err)
	}

	/* delete */
	b.Repository.Delete(blog)
}

func (b *BlogInteractor) Update(request BlogUpdateRequest) {
	/* input data -> madel object */
	blogId, err := model.NewBlogId(request.BlogId)
	if err != nil {
		b.OutputPort.RespondErorr(err)
	}

	content, err := model.NewBlogContent(request.Content)
	if err != nil {
		b.OutputPort.RespondErorr(err)
	}

	abstract, err := model.NewBlogAbstract(request.Abstract)
	if err != nil {
		b.OutputPort.RespondErorr(err)
	}

	title, err := model.NewBlogTitle(request.Title)
	if err != nil {
		b.OutputPort.RespondErorr(err)
	}

	evaluation, err := model.NewBlogEvaluation(request.Evaluation)
	if err != nil {
		b.OutputPort.RespondErorr(err)
	}

	/* find */
	exists, blog, err := b.Repository.FindById(blogId)
	if !exists {
		b.OutputPort.RespondErorr(errors.New("blog of the id dosen't exist"))
	}
	if err != nil {
		b.OutputPort.RespondErorr(err)
	}

	/* update & save */
	blog.Update(content, title, abstract, evaluation)
	b.Repository.Update(blog)
}

func (b *BlogInteractor) Register(request BlogRegisterRequest) {
	/* input data -> model object */
	userId, err := model.NewUserId(request.UserId)
	if err != nil {
		b.OutputPort.RespondErorr(err)
	}

	content, err := model.NewBlogContent(request.Content)
	if err != nil {
		b.OutputPort.RespondErorr(err)
	}

	title, err := model.NewBlogTitle(request.Title)
	if err != nil {
		b.OutputPort.RespondErorr(err)
	}

	abstract, err := model.NewBlogAbstract(request.Abstract)
	if err != nil {
		b.OutputPort.RespondErorr(err)
	}

	evaluation, err := model.NewBlogEvaluation(request.Evaluation)
	if err != nil {
		b.OutputPort.RespondErorr(err)
	}

	/* gen blog_id */
	blogService := service.Blog{}
	id, err := blogService.GenULID()
	if err != nil {
		b.OutputPort.RespondErorr(err)
	}

	blogId, err := model.NewBlogId(id.String())
	if err != nil {
		b.OutputPort.RespondErorr(err)
	}

	/* create & save */
	blog := model.NewBlog(blogId, userId, content, title, abstract, evaluation)
	b.Repository.Register(blog)
}
