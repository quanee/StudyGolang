package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Blog holds the schema definition for the Blog entity.
type Blog struct {
	ent.Schema
}

// Fields of the Blog.
func (Blog) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive(),
		field.String("title").Default("Untitled"),
		field.String("content").Default(""),
	}
}

// Edges of the Blog.
func (Blog) Edges() []ent.Edge {
	return nil
}
