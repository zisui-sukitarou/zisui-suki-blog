// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"zisui-suki-blog/ent/blog"
	"zisui-suki-blog/ent/user"

	"entgo.io/ent/dialect/sql"
)

// Blog is the model entity for the Blog schema.
type Blog struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Abstract holds the value of the "abstract" field.
	Abstract string `json:"abstract,omitempty"`
	// Evaluation holds the value of the "evaluation" field.
	Evaluation uint `json:"evaluation,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BlogQuery when eager-loading is set.
	Edges BlogEdges `json:"edges"`
}

// BlogEdges holds the relations/edges for other nodes in the graph.
type BlogEdges struct {
	// Tags holds the value of the tags edge.
	Tags []*Tag `json:"tags,omitempty"`
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// Writer holds the value of the writer edge.
	Writer *User `json:"writer,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// TagsOrErr returns the Tags value or an error if the edge
// was not loaded in eager-loading.
func (e BlogEdges) TagsOrErr() ([]*Tag, error) {
	if e.loadedTypes[0] {
		return e.Tags, nil
	}
	return nil, &NotLoadedError{edge: "tags"}
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e BlogEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// WriterOrErr returns the Writer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BlogEdges) WriterOrErr() (*User, error) {
	if e.loadedTypes[2] {
		if e.Writer == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Writer, nil
	}
	return nil, &NotLoadedError{edge: "writer"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Blog) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case blog.FieldEvaluation:
			values[i] = new(sql.NullInt64)
		case blog.FieldID, blog.FieldUserID, blog.FieldContent, blog.FieldTitle, blog.FieldAbstract:
			values[i] = new(sql.NullString)
		case blog.FieldCreatedAt, blog.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Blog", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Blog fields.
func (b *Blog) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case blog.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				b.ID = value.String
			}
		case blog.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				b.UserID = value.String
			}
		case blog.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				b.Content = value.String
			}
		case blog.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				b.Title = value.String
			}
		case blog.FieldAbstract:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field abstract", values[i])
			} else if value.Valid {
				b.Abstract = value.String
			}
		case blog.FieldEvaluation:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field evaluation", values[i])
			} else if value.Valid {
				b.Evaluation = uint(value.Int64)
			}
		case blog.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				b.CreatedAt = value.Time
			}
		case blog.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				b.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryTags queries the "tags" edge of the Blog entity.
func (b *Blog) QueryTags() *TagQuery {
	return (&BlogClient{config: b.config}).QueryTags(b)
}

// QueryUsers queries the "users" edge of the Blog entity.
func (b *Blog) QueryUsers() *UserQuery {
	return (&BlogClient{config: b.config}).QueryUsers(b)
}

// QueryWriter queries the "writer" edge of the Blog entity.
func (b *Blog) QueryWriter() *UserQuery {
	return (&BlogClient{config: b.config}).QueryWriter(b)
}

// Update returns a builder for updating this Blog.
// Note that you need to call Blog.Unwrap() before calling this method if this Blog
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Blog) Update() *BlogUpdateOne {
	return (&BlogClient{config: b.config}).UpdateOne(b)
}

// Unwrap unwraps the Blog entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Blog) Unwrap() *Blog {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Blog is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Blog) String() string {
	var builder strings.Builder
	builder.WriteString("Blog(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("user_id=")
	builder.WriteString(b.UserID)
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(b.Content)
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(b.Title)
	builder.WriteString(", ")
	builder.WriteString("abstract=")
	builder.WriteString(b.Abstract)
	builder.WriteString(", ")
	builder.WriteString("evaluation=")
	builder.WriteString(fmt.Sprintf("%v", b.Evaluation))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(b.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(b.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Blogs is a parsable slice of Blog.
type Blogs []*Blog

func (b Blogs) config(cfg config) {
	for _i := range b {
		b[_i].config = cfg
	}
}
