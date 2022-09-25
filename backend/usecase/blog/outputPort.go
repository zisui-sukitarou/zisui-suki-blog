package usecase

type BlogOutputPort interface {
	RespondBlog(BlogResponse)
	RespondBlogs([]BlogResponse)
	RespondErorr(error)
}