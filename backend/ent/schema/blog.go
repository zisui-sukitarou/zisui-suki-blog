package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Blog holds the schema definition for the Blog entity.
type Blog struct {
	ent.Schema
}

// Fields of the Blog.
func (Blog) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			StorageKey("blog_id").
			SchemaType(map[string]string{
				dialect.MySQL: "char(26)",
			}).
			Unique(),
		
		field.String("user_id"),

		field.Text("content").
			SchemaType(map[string]string{
				dialect.MySQL: "text",
			}).
			NotEmpty(),

		field.Text("title").
			SchemaType(map[string]string{
				dialect.MySQL: "text",
			}).
			NotEmpty(),

		field.Text("abstract").
			SchemaType(map[string]string{
				dialect.MySQL: "text",
			}).
			NotEmpty(),

		field.Uint("evaluation").
			SchemaType(map[string]string{
				dialect.MySQL: "int unsigned",
			}),

		field.Time("created_at").
			Immutable(),

		field.Time("updated_at"),
	}
}

// Edges of the Blog.
func (Blog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tags", Tag.Type),
		edge.From("users", User.Type).
			Ref("favors"),
		edge.To("writer", User.Type).
			Field("user_id").
			Unique().
			Required(),
	}
}
