// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"teamplayer/ent/message"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Message is the model entity for the Message schema.
type Message struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// MessageUUID holds the value of the "message_uuid" field.
	MessageUUID uuid.UUID `json:"message_uuid,omitempty"`
	// Sender holds the value of the "sender" field.
	Sender string `json:"sender,omitempty"`
	// Receiver holds the value of the "receiver" field.
	Receiver string `json:"receiver,omitempty"`
	// SendDate holds the value of the "send_date" field.
	SendDate time.Time `json:"send_date,omitempty"`
	// ReceivedDate holds the value of the "received_date" field.
	ReceivedDate time.Time `json:"received_date,omitempty"`
	// Message holds the value of the "message" field.
	Message      []string `json:"message,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Message) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case message.FieldMessage:
			values[i] = new([]byte)
		case message.FieldID:
			values[i] = new(sql.NullInt64)
		case message.FieldSender, message.FieldReceiver:
			values[i] = new(sql.NullString)
		case message.FieldSendDate, message.FieldReceivedDate:
			values[i] = new(sql.NullTime)
		case message.FieldMessageUUID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Message fields.
func (m *Message) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case message.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case message.FieldMessageUUID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field message_uuid", values[i])
			} else if value != nil {
				m.MessageUUID = *value
			}
		case message.FieldSender:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sender", values[i])
			} else if value.Valid {
				m.Sender = value.String
			}
		case message.FieldReceiver:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field receiver", values[i])
			} else if value.Valid {
				m.Receiver = value.String
			}
		case message.FieldSendDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field send_date", values[i])
			} else if value.Valid {
				m.SendDate = value.Time
			}
		case message.FieldReceivedDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field received_date", values[i])
			} else if value.Valid {
				m.ReceivedDate = value.Time
			}
		case message.FieldMessage:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field message", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &m.Message); err != nil {
					return fmt.Errorf("unmarshal field message: %w", err)
				}
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Message.
// This includes values selected through modifiers, order, etc.
func (m *Message) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// Update returns a builder for updating this Message.
// Note that you need to call Message.Unwrap() before calling this method if this Message
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Message) Update() *MessageUpdateOne {
	return NewMessageClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Message entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Message) Unwrap() *Message {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Message is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Message) String() string {
	var builder strings.Builder
	builder.WriteString("Message(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("message_uuid=")
	builder.WriteString(fmt.Sprintf("%v", m.MessageUUID))
	builder.WriteString(", ")
	builder.WriteString("sender=")
	builder.WriteString(m.Sender)
	builder.WriteString(", ")
	builder.WriteString("receiver=")
	builder.WriteString(m.Receiver)
	builder.WriteString(", ")
	builder.WriteString("send_date=")
	builder.WriteString(m.SendDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("received_date=")
	builder.WriteString(m.ReceivedDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("message=")
	builder.WriteString(fmt.Sprintf("%v", m.Message))
	builder.WriteByte(')')
	return builder.String()
}

// Messages is a parsable slice of Message.
type Messages []*Message
