package schema

import "entgo.io/ent"

// Access holds the schema definition for the Access entity.
type Access struct {
	ent.Schema
}

// Fields of the Access.
func (Access) Fields() []ent.Field {
	return nil
}

// Edges of the Access.
func (Access) Edges() []ent.Edge {
	return nil
}
