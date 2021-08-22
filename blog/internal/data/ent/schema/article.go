package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Article holds the schema definition for the Article entity.
type Article struct {
	ent.Schema
}

// Fields of the Article.
func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive(),
		field.String("content").Default("Untitled"),
	}
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return nil
}
