package blogapp

import "zisui-suki-blog/usecase/apperr"

type BlogInputPort interface {
	FindById(*BlogFindByIdRequest) error
	FindByTagName(*BlogFindByTagRequest) error
	FindByUserId(*BlogFindByUserIdRequest) error
	FindByUserIdAndTagName(*BlogFindByUserIdAndTagRequest) error
	Delete(*BlogDeleteRequest) error
	Update(*BlogUpdateRequest) error
	Register(*BlogRegisterRequest) error
}

type BlogOutputPort interface {
	RespondBlog(*BlogResponse) error
	RespondBlogOverviews([]*BlogOverviewResponse) error
	RespondErorr(*apperr.ErrorResponse) error
}