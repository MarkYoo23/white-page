package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Entry holds the schema definition for the Entry entity.
// 1. go run -mod=mod entgo.io/ent/cmd/ent new Entry
// 2. go generate ./ent

type Entry struct {
	ent.Schema

	Name    string
	Surname string
	Tel     string
}

// Fields of the Entry.
func (Entry) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("unknown"),
		field.String("surname").Default("unknown"),
		field.String("tel").Default("unknown"),
	}
}

// Edges of the Entry.
func (Entry) Edges() []ent.Edge {
	return nil
}
