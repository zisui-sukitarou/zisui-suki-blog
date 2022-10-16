package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Tag holds the schema definition for the Tag entity.
type Tag struct {
	ent.Schema
}

// Fields of the Tag.
func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			StorageKey("tag_name").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(32)",
			}).
			Unique(),

		field.Time("created_at").
			Immutable(),
	}
}

// Edges of the Tag.
func (Tag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("blogs", Blog.Type).
			Ref("tags"),
	}
}
