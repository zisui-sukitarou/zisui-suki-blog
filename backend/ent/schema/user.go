package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			StorageKey("user_id").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(15)",
			}).
			Unique(),

		field.String("name").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(32)",
			}).
			NotEmpty(),

		field.String("email").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(255)",
			}).
			NotEmpty(),

		field.String("password").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(60)",
			}).
			NotEmpty(),

		field.String("icon").
			SchemaType(map[string]string{
				dialect.MySQL: "text",
			}).
			NotEmpty(),

		field.Time("created_at").
			Immutable(),

		field.Time("updated_at"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("favors", Blog.Type),
		edge.From("blogs", Blog.Type).
			Ref("writer"),
	}
}
