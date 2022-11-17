package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Blog holds the schema definition for the Blog entity.
type Draft struct {
	ent.Schema
}

// Fields of the Blog.
func (Draft) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			StorageKey("draft_id").
			SchemaType(map[string]string{
				dialect.MySQL: "char(26)",
			}).
			Unique(),

		field.String("user_id").
			SchemaType(map[string]string{
				dialect.MySQL: "char(26)",
			}),

		field.Text("content").
			SchemaType(map[string]string{
				dialect.MySQL: "text",
			}),

		field.Text("title").
			SchemaType(map[string]string{
				dialect.MySQL: "text",
			}),

		field.Text("abstract").
			SchemaType(map[string]string{
				dialect.MySQL: "text",
			}),

		field.Uint("evaluation").
			SchemaType(map[string]string{
				dialect.MySQL: "int unsigned",
			}),

		field.Uint("status").
			SchemaType(map[string]string{
				dialect.MySQL: "int unsigned",
			}),

		field.Time("created_at").
			Immutable(),

		field.Time("updated_at"),
	}
}

// Edges of the Blog.
func (Draft) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tags", Tag.Type),
	}
}
