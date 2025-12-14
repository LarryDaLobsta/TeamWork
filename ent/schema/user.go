package schema

import (
    "time"
    "github.com/google/uuid"
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
)

// will add fields like group the user has created, roles enum

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
                field.UUID("user_uuid", uuid.UUID{}).Default(uuid.New).Unique().Immutable(),
		field.Time("user_created").Default(time.Now).Immutable(),
                field.String("first_name").NotEmpty().MaxLen(50),
                field.String("last_name").NotEmpty().MaxLen(65),
                field.String("email").NotEmpty().Unique().MaxLen(254).Sensitive(),
                field.String("username").NotEmpty().Unique().MinLen(6).MaxLen(24),
                field.String("password_hash").NotEmpty().MaxLen(255).Sensitive(),
        }
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
