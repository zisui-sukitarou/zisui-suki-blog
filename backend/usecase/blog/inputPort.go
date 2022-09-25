package usecase

type BlogInputPort interface {
	FindById(BlogFindByIdRequest)
	FindByTagId(BlogFindByTagIdRequest)
	FindByUserId(BlogFindByUserIdRequest)
	Delete(BlogDeleteRequest)
	Update(BlogUpdateRequest)
	Register(BlogRegisterRequest)
}
