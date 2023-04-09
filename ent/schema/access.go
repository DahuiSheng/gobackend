package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Access holds the schema definition for the Access entity.
type Access struct {
	ent.Schema
}

// Fields of the Access.
func (Access) Fields() []ent.Field {
	return []ent.Field{
		// UUIDの場合は、別途考える
		// ハッシュテーブルを使った方が良い場合もある
		field.String("name").
			Default("unknown"),
		field.Time("check_in").
			Default(time.Now()),
		field.Time("check_out").
			Optional().
			Nillable(),
	}
}

// Edges of the Access.
func (Access) Edges() []ent.Edge {
	return nil
}
