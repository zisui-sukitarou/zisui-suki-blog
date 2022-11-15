package draftapp

import "zisui-suki-blog/usecase/apperr"

type DraftInputPort interface {
	FindById(*DraftFindByIdRequest) error
	FindByUserId(*DraftFindByUserIdRequest) error
	FindByUserName(*DraftFindByUserNameRequest) error
	Register(*DraftRegisterRequest) error
	Update(*DraftUpdateRequest) error
	Delete(*DraftDeleteRequest) error
}

type DraftOutputPort interface {
	FindById(*DraftResponse) error
	FindByUserId([]*DraftOverviewResponse) error
	FindByUserName([]*DraftOverviewResponse) error
	Register() error
	Update(*DraftResponse) error
	Delete() error
	/* err */
	RespondErorr(*apperr.ErrorResponse) error
}