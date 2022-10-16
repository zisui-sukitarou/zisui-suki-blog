// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"zisui-suki-blog/ent/blog"
	"zisui-suki-blog/ent/predicate"
	"zisui-suki-blog/ent/tag"
	"zisui-suki-blog/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BlogUpdate is the builder for updating Blog entities.
type BlogUpdate struct {
	config
	hooks    []Hook
	mutation *BlogMutation
}

// Where appends a list predicates to the BlogUpdate builder.
func (bu *BlogUpdate) Where(ps ...predicate.Blog) *BlogUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetUserID sets the "user_id" field.
func (bu *BlogUpdate) SetUserID(s string) *BlogUpdate {
	bu.mutation.SetUserID(s)
	return bu
}

// SetContent sets the "content" field.
func (bu *BlogUpdate) SetContent(s string) *BlogUpdate {
	bu.mutation.SetContent(s)
	return bu
}

// SetTitle sets the "title" field.
func (bu *BlogUpdate) SetTitle(s string) *BlogUpdate {
	bu.mutation.SetTitle(s)
	return bu
}

// SetAbstract sets the "abstract" field.
func (bu *BlogUpdate) SetAbstract(s string) *BlogUpdate {
	bu.mutation.SetAbstract(s)
	return bu
}

// SetEvaluation sets the "evaluation" field.
func (bu *BlogUpdate) SetEvaluation(u uint) *BlogUpdate {
	bu.mutation.ResetEvaluation()
	bu.mutation.SetEvaluation(u)
	return bu
}

// AddEvaluation adds u to the "evaluation" field.
func (bu *BlogUpdate) AddEvaluation(u int) *BlogUpdate {
	bu.mutation.AddEvaluation(u)
	return bu
}

// SetUpdatedAt sets the "updated_at" field.
func (bu *BlogUpdate) SetUpdatedAt(t time.Time) *BlogUpdate {
	bu.mutation.SetUpdatedAt(t)
	return bu
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (bu *BlogUpdate) AddTagIDs(ids ...string) *BlogUpdate {
	bu.mutation.AddTagIDs(ids...)
	return bu
}

// AddTags adds the "tags" edges to the Tag entity.
func (bu *BlogUpdate) AddTags(t ...*Tag) *BlogUpdate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bu.AddTagIDs(ids...)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (bu *BlogUpdate) AddUserIDs(ids ...string) *BlogUpdate {
	bu.mutation.AddUserIDs(ids...)
	return bu
}

// AddUsers adds the "users" edges to the User entity.
func (bu *BlogUpdate) AddUsers(u ...*User) *BlogUpdate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return bu.AddUserIDs(ids...)
}

// SetWriterID sets the "writer" edge to the User entity by ID.
func (bu *BlogUpdate) SetWriterID(id string) *BlogUpdate {
	bu.mutation.SetWriterID(id)
	return bu
}

// SetWriter sets the "writer" edge to the User entity.
func (bu *BlogUpdate) SetWriter(u *User) *BlogUpdate {
	return bu.SetWriterID(u.ID)
}

// Mutation returns the BlogMutation object of the builder.
func (bu *BlogUpdate) Mutation() *BlogMutation {
	return bu.mutation
}

// ClearTags clears all "tags" edges to the Tag entity.
func (bu *BlogUpdate) ClearTags() *BlogUpdate {
	bu.mutation.ClearTags()
	return bu
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (bu *BlogUpdate) RemoveTagIDs(ids ...string) *BlogUpdate {
	bu.mutation.RemoveTagIDs(ids...)
	return bu
}

// RemoveTags removes "tags" edges to Tag entities.
func (bu *BlogUpdate) RemoveTags(t ...*Tag) *BlogUpdate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bu.RemoveTagIDs(ids...)
}

// ClearUsers clears all "users" edges to the User entity.
func (bu *BlogUpdate) ClearUsers() *BlogUpdate {
	bu.mutation.ClearUsers()
	return bu
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (bu *BlogUpdate) RemoveUserIDs(ids ...string) *BlogUpdate {
	bu.mutation.RemoveUserIDs(ids...)
	return bu
}

// RemoveUsers removes "users" edges to User entities.
func (bu *BlogUpdate) RemoveUsers(u ...*User) *BlogUpdate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return bu.RemoveUserIDs(ids...)
}

// ClearWriter clears the "writer" edge to the User entity.
func (bu *BlogUpdate) ClearWriter() *BlogUpdate {
	bu.mutation.ClearWriter()
	return bu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BlogUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bu.hooks) == 0 {
		if err = bu.check(); err != nil {
			return 0, err
		}
		affected, err = bu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BlogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bu.check(); err != nil {
				return 0, err
			}
			bu.mutation = mutation
			affected, err = bu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bu.hooks) - 1; i >= 0; i-- {
			if bu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BlogUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BlogUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BlogUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bu *BlogUpdate) check() error {
	if v, ok := bu.mutation.Content(); ok {
		if err := blog.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Blog.content": %w`, err)}
		}
	}
	if v, ok := bu.mutation.Title(); ok {
		if err := blog.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Blog.title": %w`, err)}
		}
	}
	if v, ok := bu.mutation.Abstract(); ok {
		if err := blog.AbstractValidator(v); err != nil {
			return &ValidationError{Name: "abstract", err: fmt.Errorf(`ent: validator failed for field "Blog.abstract": %w`, err)}
		}
	}
	if _, ok := bu.mutation.WriterID(); bu.mutation.WriterCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Blog.writer"`)
	}
	return nil
}

func (bu *BlogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   blog.Table,
			Columns: blog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: blog.FieldID,
			},
		},
	}
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: blog.FieldContent,
		})
	}
	if value, ok := bu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: blog.FieldTitle,
		})
	}
	if value, ok := bu.mutation.Abstract(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: blog.FieldAbstract,
		})
	}
	if value, ok := bu.mutation.Evaluation(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: blog.FieldEvaluation,
		})
	}
	if value, ok := bu.mutation.AddedEvaluation(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: blog.FieldEvaluation,
		})
	}
	if value, ok := bu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: blog.FieldUpdatedAt,
		})
	}
	if bu.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   blog.TagsTable,
			Columns: blog.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: tag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedTagsIDs(); len(nodes) > 0 && !bu.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   blog.TagsTable,
			Columns: blog.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   blog.TagsTable,
			Columns: blog.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   blog.UsersTable,
			Columns: blog.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedUsersIDs(); len(nodes) > 0 && !bu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   blog.UsersTable,
			Columns: blog.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   blog.UsersTable,
			Columns: blog.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.WriterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   blog.WriterTable,
			Columns: []string{blog.WriterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.WriterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   blog.WriterTable,
			Columns: []string{blog.WriterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{blog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// BlogUpdateOne is the builder for updating a single Blog entity.
type BlogUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BlogMutation
}

// SetUserID sets the "user_id" field.
func (buo *BlogUpdateOne) SetUserID(s string) *BlogUpdateOne {
	buo.mutation.SetUserID(s)
	return buo
}

// SetContent sets the "content" field.
func (buo *BlogUpdateOne) SetContent(s string) *BlogUpdateOne {
	buo.mutation.SetContent(s)
	return buo
}

// SetTitle sets the "title" field.
func (buo *BlogUpdateOne) SetTitle(s string) *BlogUpdateOne {
	buo.mutation.SetTitle(s)
	return buo
}

// SetAbstract sets the "abstract" field.
func (buo *BlogUpdateOne) SetAbstract(s string) *BlogUpdateOne {
	buo.mutation.SetAbstract(s)
	return buo
}

// SetEvaluation sets the "evaluation" field.
func (buo *BlogUpdateOne) SetEvaluation(u uint) *BlogUpdateOne {
	buo.mutation.ResetEvaluation()
	buo.mutation.SetEvaluation(u)
	return buo
}

// AddEvaluation adds u to the "evaluation" field.
func (buo *BlogUpdateOne) AddEvaluation(u int) *BlogUpdateOne {
	buo.mutation.AddEvaluation(u)
	return buo
}

// SetUpdatedAt sets the "updated_at" field.
func (buo *BlogUpdateOne) SetUpdatedAt(t time.Time) *BlogUpdateOne {
	buo.mutation.SetUpdatedAt(t)
	return buo
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (buo *BlogUpdateOne) AddTagIDs(ids ...string) *BlogUpdateOne {
	buo.mutation.AddTagIDs(ids...)
	return buo
}

// AddTags adds the "tags" edges to the Tag entity.
func (buo *BlogUpdateOne) AddTags(t ...*Tag) *BlogUpdateOne {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return buo.AddTagIDs(ids...)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (buo *BlogUpdateOne) AddUserIDs(ids ...string) *BlogUpdateOne {
	buo.mutation.AddUserIDs(ids...)
	return buo
}

// AddUsers adds the "users" edges to the User entity.
func (buo *BlogUpdateOne) AddUsers(u ...*User) *BlogUpdateOne {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return buo.AddUserIDs(ids...)
}

// SetWriterID sets the "writer" edge to the User entity by ID.
func (buo *BlogUpdateOne) SetWriterID(id string) *BlogUpdateOne {
	buo.mutation.SetWriterID(id)
	return buo
}

// SetWriter sets the "writer" edge to the User entity.
func (buo *BlogUpdateOne) SetWriter(u *User) *BlogUpdateOne {
	return buo.SetWriterID(u.ID)
}

// Mutation returns the BlogMutation object of the builder.
func (buo *BlogUpdateOne) Mutation() *BlogMutation {
	return buo.mutation
}

// ClearTags clears all "tags" edges to the Tag entity.
func (buo *BlogUpdateOne) ClearTags() *BlogUpdateOne {
	buo.mutation.ClearTags()
	return buo
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (buo *BlogUpdateOne) RemoveTagIDs(ids ...string) *BlogUpdateOne {
	buo.mutation.RemoveTagIDs(ids...)
	return buo
}

// RemoveTags removes "tags" edges to Tag entities.
func (buo *BlogUpdateOne) RemoveTags(t ...*Tag) *BlogUpdateOne {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return buo.RemoveTagIDs(ids...)
}

// ClearUsers clears all "users" edges to the User entity.
func (buo *BlogUpdateOne) ClearUsers() *BlogUpdateOne {
	buo.mutation.ClearUsers()
	return buo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (buo *BlogUpdateOne) RemoveUserIDs(ids ...string) *BlogUpdateOne {
	buo.mutation.RemoveUserIDs(ids...)
	return buo
}

// RemoveUsers removes "users" edges to User entities.
func (buo *BlogUpdateOne) RemoveUsers(u ...*User) *BlogUpdateOne {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return buo.RemoveUserIDs(ids...)
}

// ClearWriter clears the "writer" edge to the User entity.
func (buo *BlogUpdateOne) ClearWriter() *BlogUpdateOne {
	buo.mutation.ClearWriter()
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BlogUpdateOne) Select(field string, fields ...string) *BlogUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Blog entity.
func (buo *BlogUpdateOne) Save(ctx context.Context) (*Blog, error) {
	var (
		err  error
		node *Blog
	)
	if len(buo.hooks) == 0 {
		if err = buo.check(); err != nil {
			return nil, err
		}
		node, err = buo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BlogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = buo.check(); err != nil {
				return nil, err
			}
			buo.mutation = mutation
			node, err = buo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(buo.hooks) - 1; i >= 0; i-- {
			if buo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = buo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, buo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Blog)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from BlogMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BlogUpdateOne) SaveX(ctx context.Context) *Blog {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BlogUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BlogUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (buo *BlogUpdateOne) check() error {
	if v, ok := buo.mutation.Content(); ok {
		if err := blog.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Blog.content": %w`, err)}
		}
	}
	if v, ok := buo.mutation.Title(); ok {
		if err := blog.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Blog.title": %w`, err)}
		}
	}
	if v, ok := buo.mutation.Abstract(); ok {
		if err := blog.AbstractValidator(v); err != nil {
			return &ValidationError{Name: "abstract", err: fmt.Errorf(`ent: validator failed for field "Blog.abstract": %w`, err)}
		}
	}
	if _, ok := buo.mutation.WriterID(); buo.mutation.WriterCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Blog.writer"`)
	}
	return nil
}

func (buo *BlogUpdateOne) sqlSave(ctx context.Context) (_node *Blog, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   blog.Table,
			Columns: blog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: blog.FieldID,
			},
		},
	}
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Blog.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, blog.FieldID)
		for _, f := range fields {
			if !blog.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != blog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: blog.FieldContent,
		})
	}
	if value, ok := buo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: blog.FieldTitle,
		})
	}
	if value, ok := buo.mutation.Abstract(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: blog.FieldAbstract,
		})
	}
	if value, ok := buo.mutation.Evaluation(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: blog.FieldEvaluation,
		})
	}
	if value, ok := buo.mutation.AddedEvaluation(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint,
			Value:  value,
			Column: blog.FieldEvaluation,
		})
	}
	if value, ok := buo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: blog.FieldUpdatedAt,
		})
	}
	if buo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   blog.TagsTable,
			Columns: blog.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: tag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedTagsIDs(); len(nodes) > 0 && !buo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   blog.TagsTable,
			Columns: blog.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   blog.TagsTable,
			Columns: blog.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   blog.UsersTable,
			Columns: blog.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !buo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   blog.UsersTable,
			Columns: blog.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   blog.UsersTable,
			Columns: blog.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.WriterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   blog.WriterTable,
			Columns: []string{blog.WriterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.WriterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   blog.WriterTable,
			Columns: []string{blog.WriterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Blog{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{blog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
