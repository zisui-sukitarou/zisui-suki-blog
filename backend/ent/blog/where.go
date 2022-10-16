// Code generated by ent, DO NOT EDIT.

package blog

import (
	"time"
	"zisui-suki-blog/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// Abstract applies equality check predicate on the "abstract" field. It's identical to AbstractEQ.
func Abstract(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAbstract), v))
	})
}

// Evaluation applies equality check predicate on the "evaluation" field. It's identical to EvaluationEQ.
func Evaluation(v uint) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEvaluation), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.Blog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.Blog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	})
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	})
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	})
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	})
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUserID), v))
	})
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUserID), v))
	})
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUserID), v))
	})
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUserID), v))
	})
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUserID), v))
	})
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContent), v))
	})
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Blog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldContent), v...))
	})
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Blog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldContent), v...))
	})
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContent), v))
	})
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContent), v))
	})
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContent), v))
	})
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContent), v))
	})
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldContent), v))
	})
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldContent), v))
	})
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldContent), v))
	})
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldContent), v))
	})
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldContent), v))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Blog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Blog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// AbstractEQ applies the EQ predicate on the "abstract" field.
func AbstractEQ(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAbstract), v))
	})
}

// AbstractNEQ applies the NEQ predicate on the "abstract" field.
func AbstractNEQ(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAbstract), v))
	})
}

// AbstractIn applies the In predicate on the "abstract" field.
func AbstractIn(vs ...string) predicate.Blog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAbstract), v...))
	})
}

// AbstractNotIn applies the NotIn predicate on the "abstract" field.
func AbstractNotIn(vs ...string) predicate.Blog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAbstract), v...))
	})
}

// AbstractGT applies the GT predicate on the "abstract" field.
func AbstractGT(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAbstract), v))
	})
}

// AbstractGTE applies the GTE predicate on the "abstract" field.
func AbstractGTE(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAbstract), v))
	})
}

// AbstractLT applies the LT predicate on the "abstract" field.
func AbstractLT(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAbstract), v))
	})
}

// AbstractLTE applies the LTE predicate on the "abstract" field.
func AbstractLTE(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAbstract), v))
	})
}

// AbstractContains applies the Contains predicate on the "abstract" field.
func AbstractContains(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAbstract), v))
	})
}

// AbstractHasPrefix applies the HasPrefix predicate on the "abstract" field.
func AbstractHasPrefix(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAbstract), v))
	})
}

// AbstractHasSuffix applies the HasSuffix predicate on the "abstract" field.
func AbstractHasSuffix(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAbstract), v))
	})
}

// AbstractEqualFold applies the EqualFold predicate on the "abstract" field.
func AbstractEqualFold(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAbstract), v))
	})
}

// AbstractContainsFold applies the ContainsFold predicate on the "abstract" field.
func AbstractContainsFold(v string) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAbstract), v))
	})
}

// EvaluationEQ applies the EQ predicate on the "evaluation" field.
func EvaluationEQ(v uint) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEvaluation), v))
	})
}

// EvaluationNEQ applies the NEQ predicate on the "evaluation" field.
func EvaluationNEQ(v uint) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEvaluation), v))
	})
}

// EvaluationIn applies the In predicate on the "evaluation" field.
func EvaluationIn(vs ...uint) predicate.Blog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEvaluation), v...))
	})
}

// EvaluationNotIn applies the NotIn predicate on the "evaluation" field.
func EvaluationNotIn(vs ...uint) predicate.Blog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEvaluation), v...))
	})
}

// EvaluationGT applies the GT predicate on the "evaluation" field.
func EvaluationGT(v uint) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEvaluation), v))
	})
}

// EvaluationGTE applies the GTE predicate on the "evaluation" field.
func EvaluationGTE(v uint) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEvaluation), v))
	})
}

// EvaluationLT applies the LT predicate on the "evaluation" field.
func EvaluationLT(v uint) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEvaluation), v))
	})
}

// EvaluationLTE applies the LTE predicate on the "evaluation" field.
func EvaluationLTE(v uint) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEvaluation), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Blog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Blog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Blog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Blog {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// HasTags applies the HasEdge predicate on the "tags" edge.
func HasTags() predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TagsTable, TagFieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, TagsTable, TagsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTagsWith applies the HasEdge predicate on the "tags" edge with a given conditions (other predicates).
func HasTagsWith(preds ...predicate.Tag) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TagsInverseTable, TagFieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, TagsTable, TagsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UsersTable, UserFieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UsersTable, UsersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UsersInverseTable, UserFieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UsersTable, UsersPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasWriter applies the HasEdge predicate on the "writer" edge.
func HasWriter() predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WriterTable, UserFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, WriterTable, WriterColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWriterWith applies the HasEdge predicate on the "writer" edge with a given conditions (other predicates).
func HasWriterWith(preds ...predicate.User) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WriterInverseTable, UserFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, WriterTable, WriterColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Blog) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Blog) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Blog) predicate.Blog {
	return predicate.Blog(func(s *sql.Selector) {
		p(s.Not())
	})
}
