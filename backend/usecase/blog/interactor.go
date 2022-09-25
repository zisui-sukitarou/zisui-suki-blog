package usecase

import (
	"cook-blog/domain/model"
	"cook-blog/domain/repository"
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
	blogId, err := model.NewBlogId(request.Id)
	if err != nil {
		b.OutputPort.RespondErorr(err)
	}

	/* response */
	blog := b.Repository.FindById(blogId)
	response := NewBlogResponse(blog)
	b.OutputPort.RespondBlog(response)
}

func (b *BlogInteractor) FindByUserId(request BlogFindByUserIdRequest) {
	/* input data -> model object */
	userId, err := model.NewUserId(request.UserId)
	if err != nil {
		b.OutputPort.RespondErorr(err)
	}

	/* response */
	blogs := b.Repository.FindByUserId(userId)
	response := NewBlogsResponse(blogs)
	b.OutputPort.RespondBlogs(response)
}

func (b *BlogInteractor) FindByTagId(request BlogFindByTagIdRequest) {

}

func (b *BlogInteractor) Delete(request BlogDeleteRequest) {
	/* input data -> model object */
	blogId, err := model.NewBlogId(request.Id)
	if err != nil {
		b.OutputPort.RespondErorr(err)
	}

	/* delete */
	blog := b.Repository.FindById(blogId)
	b.Repository.Delete(blog)
}

func (b *BlogInteractor) Update(request BlogUpdateRequest) {
	/* input data -> madel object */
	blogId, err := model.NewBlogId(request.Id)
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

	/* update */
	blog := b.Repository.FindById(blogId)
	blog.Update(model.CommandUpdateBlog{
		Content:  content,
		Abstract: abstract,
		Title:    title,
	})
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

	/* register */
	blog, err := model.NewBlog(model.CommandNewBlog{
		UserId:   userId,
		Content:  content,
		Title:    title,
		Abstract: abstract,
	})
	b.Repository.Register(blog)
}
