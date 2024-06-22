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
                field.UUID("user_uuid", uuid.UUID{}).Default(uuid.New),
                field.Int("id"),
		field.Time("user_created").Default(time.Now),
                field.String("first_name").NotEmpty().MaxLen(30),
                field.String("last_name").NotEmpty().MaxLen(30),
                field.String("username").NotEmpty().Unique().Sensitive(),
                field.String("password").NotEmpty().Unique().MinLen(10).Sensitive(),
        }

}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
