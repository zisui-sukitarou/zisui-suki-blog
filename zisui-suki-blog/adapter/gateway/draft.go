package gateway

import (
	"context"
	"zisui-suki-blog/domain/model"
	"zisui-suki-blog/domain/repository"
	"zisui-suki-blog/ent"
	"zisui-suki-blog/ent/draft"
)

type DraftGateway struct {
	Client  *ent.Client
	Context *context.Context
}

func NewDraftGateway(
	client *ent.Client,
	ctx *context.Context,
) repository.DraftRepository {
	return &DraftGateway{
		Client:  client,
		Context: ctx,
	}
}

func (g *DraftGateway) FindById(draftId model.DraftId) (bool, *repository.DraftData, error) {
	draft, err := g.Client.Debug().Draft.
		Query().
		Where(draft.IDEQ(string(draftId))).
		WithTags().
		First(*g.Context)
	if err != nil {
		return false, &repository.DraftData{}, err
	}

	/* create draft model object */
	b := model.NewDraft(
		model.DraftId(draft.ID),
		model.UserId(draft.UserID),
		model.DraftContent(draft.Content),
		model.DraftTitle(draft.Title),
		model.DraftAbstract(draft.Abstract),
		model.DraftEvaluation(draft.Evaluation),
		model.DraftStatus(draft.Status),
		draft.CreatedAt,
		draft.UpdatedAt,
	)

	/* create tags model object */
	var tags []*model.Tag
	for _, tag := range draft.Edges.Tags {
		tags = append(tags, model.NewTag(
			model.TagName(tag.ID),
			model.TagIcon(tag.Icon),
			model.TagStatus(tag.Status),
			tag.CreatedAt,
			tag.UpdatedAt,
		))
	}

	return true, repository.NewDraftData(b, tags), nil
}

func (g *DraftGateway) FindByUserId(userId model.UserId, begin uint, end uint) ([]*repository.DraftData, error) {

	// blog.content 以外を取得
	drafts, err := g.Client.Debug().Draft.Query().
		Select(draft.FieldID).
		Select(draft.FieldUserID).
		Select(draft.FieldTitle).
		Select(draft.FieldAbstract).
		Select(draft.FieldEvaluation).
		Select(draft.FieldStatus).
		Select(draft.FieldCreatedAt).
		Select(draft.FieldUpdatedAt).
		Where(draft.UserIDEQ(string(userId))).
		WithTags().
		Offset(int(begin)).
		Limit(int(end)).
		All(*g.Context)
	if err != nil {
		return []*repository.DraftData{}, err
	}

	var overviews []*repository.DraftData
	for _, draft := range drafts {
		/* create draft model object */
		d := model.NewDraft(
			model.DraftId(draft.ID),
			model.UserId(draft.UserID),
			model.DraftContent(draft.Content),
			model.DraftTitle(draft.Title),
			model.DraftAbstract(draft.Abstract),
			model.DraftEvaluation(draft.Evaluation),
			model.DraftStatus(draft.Status),
			draft.CreatedAt,
			draft.UpdatedAt,
		)

		/* create tags model object */
		var tags []*model.Tag
		for _, tag := range draft.Edges.Tags {
			tags = append(tags, model.NewTag(
				model.TagName(tag.ID),
				model.TagIcon(tag.Icon),
				model.TagStatus(tag.Status),
				tag.CreatedAt,
				tag.UpdatedAt,
			))
		}

		/* create draft overviews */
		overviews = append(overviews, repository.NewDraftData(d, tags))
	}
	return overviews, nil
}

func (g *DraftGateway) Register(draft *model.Draft) error {
	_, err := g.Client.Debug().Draft.Create().
		SetID(string(draft.DraftId)).
		SetUserID(string(draft.UserId)).
		SetTitle(string(draft.Title)).
		SetAbstract(string(draft.Abstract)).
		SetContent(string(draft.Content)).
		SetEvaluation(uint(draft.Evaluation)).
		SetStatus(uint(draft.Status)).
		SetCreatedAt(draft.CreatedAt).
		SetUpdatedAt(draft.UpdatedAt).
		Save(*g.Context)
	return err
}

func (g *DraftGateway) RegisterTags(draftId model.DraftId, tagNames []model.TagName) error {
	/* tagNames ([]model.TagName) -> tagNameStrings ([]string) */
	var tagNameStrings []string
	for _, v := range tagNames {
		tagNameStrings = append(tagNameStrings, string(v))
	}

	/* save */
	_, err := g.Client.Debug().Draft.Update().
		Where(draft.IDEQ(string(draftId))).
		AddTagIDs(tagNameStrings...).
		Save(*g.Context)
	return err
}

func (g *DraftGateway) Update(draft *model.Draft) error {
	/* save */
	_, err := g.Client.Debug().Draft.
		UpdateOneID(string(draft.DraftId)).
		SetAbstract(string(draft.Abstract)).
		SetContent(string(draft.Content)).
		SetTitle(string(draft.Title)).
		SetEvaluation(uint(draft.Evaluation)).
		Save(*g.Context)
	return err
}

func (g *DraftGateway) UpdateTags(draftId model.DraftId, updatingTagNames []model.TagName) error {
	/* tagNames ([]model.TagName) -> tagNameStrings ([]string) */
	var updatingTagNameStrings []string
	for _, v := range updatingTagNames {
		updatingTagNameStrings = append(updatingTagNameStrings, string(v))
	}

	/* delete */
	err := g.Client.Debug().Draft.UpdateOneID(string(draftId)).ClearTags().Exec(*g.Context)
	if err != nil {
		return err
	}

	/* save */
	return g.RegisterTags(draftId, updatingTagNames)
}

func (g *DraftGateway) Delete(draft *model.Draft) error {
	return nil
}
