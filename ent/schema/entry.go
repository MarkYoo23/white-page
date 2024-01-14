package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Entry holds the schema definition for the Entry entity.
// 1. go run -mod=mod entgo.io/ent/cmd/ent new Entry
// 1. go run -mod=mod entgo.io/ent/cmd/ent new User

// 2. go generate ./ent

type Entry struct {
	ent.Schema

	Name     string
	Surname  string
	Tel      string
	CreateAt time.Time
}

// Fields of the Entry.
func (Entry) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("unknown"),
		field.String("surname").Default("unknown"),
		field.String("tel").Unique(),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Entry.
func (Entry) Edges() []ent.Edge {
	return nil
}
