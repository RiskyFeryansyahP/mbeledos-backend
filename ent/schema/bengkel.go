package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// Bengkel holds the schema definition for the Bengkel entity.
type Bengkel struct {
	ent.Schema
}

// Fields of the Bengkel.
func (Bengkel) Fields() []ent.Field {
	return []ent.Field{
		field.String("kode_bengkel").
			Unique(),
		field.String("nama_bengkel").
			NotEmpty().
			Default("Bengkel Unknown"),
		field.String("alamat_bengkel").
			NotEmpty(),
		field.Float("latitude"),
		field.Float("longitude"),
	}
}

// Edges of the Bengkel.
func (Bengkel) Edges() []ent.Edge {
	return nil
}
