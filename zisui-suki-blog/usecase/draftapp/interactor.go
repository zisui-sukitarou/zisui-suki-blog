package draftapp

import (
	"errors"
	"time"
	"zisui-suki-blog/domain/model"
	"zisui-suki-blog/domain/repository"
	"zisui-suki-blog/domain/service"
	"zisui-suki-blog/usecase/apperr"
)

type DraftInteractor struct {
	OutputPort DraftOutputPort
	DraftRepo  repository.DraftRepository
	UserRepo   repository.UserRepository
	TagRepo    repository.TagRepository
}

/* constructor */
func NewDraftInteractor(
	outputPort DraftOutputPort,
	draftRepo repository.DraftRepository,
	userRepo repository.UserRepository,
	tagRepo repository.TagRepository,
) DraftInputPort {
	return &DraftInteractor{
		OutputPort: outputPort,
		DraftRepo:  draftRepo,
		UserRepo:   userRepo,
		TagRepo:    tagRepo,
	}
}

func (d *DraftInteractor) FindById(request *DraftFindByIdRequest) error {
	/* input data ->　model object */
	draftId, err := model.NewDraftId(request.DraftId)
	if err != nil {
		return d.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* find draft by id */
	exists, draft, err := d.DraftRepo.FindById(draftId)
	if !exists {
		return d.OutputPort.RespondErorr(apperr.NewErrorResponse(errors.New("draft of the id dosen't exist")))
	}
	if err != nil {
		return d.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* find user by user_id */
	exists, user, err := d.UserRepo.FindById(draft.Draft.UserId)
	if !exists {
		return d.OutputPort.RespondErorr(apperr.NewErrorResponse(errors.New("user of the id dosen't exist")))
	}
	if err != nil {
		return d.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* response */
	return d.OutputPort.FindById(NewDraftResponse(draft, user))
}

func (d *DraftInteractor) FindByUserId(request *DraftFindByUserIdRequest) error {
	/* input data -> model object */
	userId, err := model.NewUserId(request.UserId)
	if err != nil {
		return d.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* find by user id */
	drafts, err := d.DraftRepo.FindByUserId(userId, request.Begin, request.End)
	if err != nil {
		return d.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	var response []*DraftOverviewResponse
	for _, draft := range drafts {
		/* find user by user_id */
		exists, user, err := d.UserRepo.FindById(draft.Draft.UserId)
		if !exists {
			return d.OutputPort.RespondErorr(apperr.NewErrorResponse(errors.New("user of the id dosen't exist")))
		}
		if err != nil {
			return d.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
		}

		response = append(response, NewDraftOverviewResponse(draft, user))
	}

	/* response */
	return d.OutputPort.FindByUserId(response)
}

func (d *DraftInteractor) FindByUserName(request *DraftFindByUserNameRequest) error {
	/* input data -> model object */
	userName, err := model.NewUserName(request.UserName)
	if err != nil {
		return d.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* userName -> userId */
	exists, user, err := d.UserRepo.FindByUserName(userName)
	if !exists || err != nil {
		return d.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* find by user id */
	drafts, err := d.DraftRepo.FindByUserId(user.UserId, request.Begin, request.End)
	if err != nil {
		return d.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	var response []*DraftOverviewResponse
	for _, draft := range drafts {
		/* find user by user_id */
		exists, user, err := d.UserRepo.FindById(draft.Draft.UserId)
		if !exists {
			return d.OutputPort.RespondErorr(apperr.NewErrorResponse(errors.New("user of the id dosen't exist")))
		}
		if err != nil {
			return d.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
		}

		response = append(response, NewDraftOverviewResponse(draft, user))
	}

	/* response */
	return d.OutputPort.FindByUserName(response)
}

func (d *DraftInteractor) Delete(request *DraftDeleteRequest) error {
	/* input data -> model object */
	draftId, err := model.NewDraftId(request.DraftId)
	if err != nil {
		return d.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* find by id */
	exists, draft, err := d.DraftRepo.FindById(draftId)
	if err != nil {
		d.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}
	if !exists {
		d.OutputPort.RespondErorr(apperr.NewErrorResponse(errors.New("draft of the id dosen't exist")))
	}

	/* delete */
	return d.DraftRepo.Delete(draft.Draft)
}

func (d *DraftInteractor) Update(request *DraftUpdateRequest) error {
	/* input data -> madel object */
	draftId, err := model.NewDraftId(request.DraftId)
	if err != nil {
		return d.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	content, err := model.NewDraftContent(request.Content)
	if err != nil {
		return d.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	abstract, err := model.NewDraftAbstract(request.Abstract)
	if err != nil {
		return d.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	title, err := model.NewDraftTitle(request.Title)
	if err != nil {
		return d.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	evaluation, err := model.NewDraftEvaluation(request.Evaluation)
	if err != nil {
		return d.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	var tagNames []model.TagName
	for _, tag := range request.Tags {
		tagName, err := model.NewTagName(tag)
		if err != nil {
			return d.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
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
		d.TagRepo.Register(tag)
	}

	/* find draft by id */
	exists, draft, err := d.DraftRepo.FindById(draftId)
	if !exists {
		return d.OutputPort.RespondErorr(apperr.NewErrorResponse(errors.New("draft of the id dosen't exist")))
	}
	if err != nil {
		return d.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* update draft & save */
	draft.Draft.Update(content, title, abstract, evaluation)
	err = d.DraftRepo.Update(draft.Draft)
	if err != nil {
		return d.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* update tags */
	err = d.DraftRepo.UpdateTags(draftId, tagNames)
	if err != nil {
		return d.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* find user by user_id */
	exists, user, err := d.UserRepo.FindById(draft.Draft.UserId)
	if !exists {
		return d.OutputPort.RespondErorr(apperr.NewErrorResponse(errors.New("user of the id dosen't exist")))
	}
	if err != nil {
		return d.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* response */
	return d.OutputPort.Update(NewDraftResponse(draft, user))
}

func (d *DraftInteractor) Register(request *DraftRegisterRequest) error {
	/* input data -> model object */
	userId, err := model.NewUserId(request.UserId)
	if err != nil {
		return d.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	content, err := model.NewDraftContent(request.Content)
	if err != nil {
		return d.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	title, err := model.NewDraftTitle(request.Title)
	if err != nil {
		return d.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	abstract, err := model.NewDraftAbstract(request.Abstract)
	if err != nil {
		return d.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	evaluation, err := model.NewDraftEvaluation(request.Evaluation)
	if err != nil {
		return d.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	status, err := model.NewDraftStatus(request.Status)
	if err != nil {
		return d.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	var tagNames []model.TagName
	for _, tag := range request.Tags {
		tagName, err := model.NewTagName(tag)
		if err != nil {
			return d.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
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
		d.TagRepo.Register(tag)
	}

	/* gen draft_id */
	id, err := d.service().GenULID()
	if err != nil {
		return d.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	draftId, err := model.NewDraftId(id.String())
	if err != nil {
		return d.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* create draft & save */
	draft := model.NewDraft(draftId, userId, content, title, abstract, evaluation, status, time.Now(), time.Now())
	err = d.DraftRepo.Register(draft)
	if err != nil {
		return d.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* register tags info */
	err = d.DraftRepo.UpdateTags(draftId, tagNames)
	if err != nil {
		return d.OutputPort.RespondErorr(apperr.NewErrorResponse(err))
	}

	/* response */
	return d.OutputPort.Register()
}

/* create draft service */
func (d *DraftInteractor) service() *service.Draft {
	return &service.Draft{}
}

/* model.TagName[] contains model.TagName ?*/
func (d *DraftInteractor) contain(tags []*model.Tag, target model.TagName) bool {
	for _, tag := range tags {
		if tag.TagName == target {
			return true
		}
	}
	return false
}
