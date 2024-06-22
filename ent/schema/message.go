package schema
import (
    "time"
    "github.com/google/uuid"
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
                field.UUID("message_uuid", uuid.UUID{}).Default(uuid.New),
                field.Int("id"),
                // need to not be empty
                field.String("sender"),
                // need to not be empty
                field.String("receiver"),

                // need to not be empty/also known as create date
                field.Time("send_date").Default(time.Now),
                // need to not be empty
                field.Time("received_date").Default(time.Now),
                // optional empty with the message just space
                field.JSON("message", []string{}).Optional(),
        }
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return nil
}
