package blogapp

type BlogOutputPort interface {
	RespondBlog(BlogResponse)
	RespondBlogs([]BlogResponse)
	RespondErorr(error)
}