package blogapp

type BlogInputPort interface {
	FindById(BlogFindByIdRequest) error
	FindByTagName(BlogFindByTagRequest) error
	FindByUserId(BlogFindByUserIdRequest) error
	Delete(BlogDeleteRequest) error
	Update(BlogUpdateRequest) error
	Register(BlogRegisterRequest) error
}

type BlogOutputPort interface {
	RespondBlog(BlogResponse) error
	RespondBlogs([]BlogResponse) error
	RespondErorr(error) error
}