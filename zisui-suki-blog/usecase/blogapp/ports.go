package blogapp

import "zisui-suki-blog/usecase/apperr"

type BlogInputPort interface {
	FindById(*BlogFindByIdRequest) error
	FindByTagName(*BlogFindByTagRequest) error
	FindByUserName(*BlogFindByUserNameRequest) error
	FindByUserNameAndTagName(*BlogFindByUserNameAndTagRequest) error
	Register(*BlogRegisterRequest) error
	Update(*BlogUpdateRequest) error
	Delete(*BlogDeleteRequest) error
}

type BlogOutputPort interface {
	FindById(*BlogResponse) error
	FindByTagName([]*BlogOverviewResponse) error
	FindByUserName([]*BlogOverviewResponse) error
	FindByUserNameAndTagName([]*BlogOverviewResponse) error
	Register() error
	Update(*BlogResponse) error
	Delete() error
	/* err */
	RespondErorr(*apperr.ErrorResponse) error
}