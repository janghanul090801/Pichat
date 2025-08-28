package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			StorageKey("NAME"),
		field.String("email").
			NotEmpty().
			StorageKey("EMAIL"),
		field.String("password").
			NotEmpty().
			Sensitive().
			StorageKey("PASSWORD"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
