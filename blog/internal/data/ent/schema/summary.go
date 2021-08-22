package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Summary holds the schema definition for the Summary entity.
type Summary struct {
	ent.Schema
}

// Fields of the Summary.
func (Summary) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive(),
		field.String("title").Default("Untitled"),
	}
}

// Edges of the Summary.
func (Summary) Edges() []ent.Edge {
	return nil
}
