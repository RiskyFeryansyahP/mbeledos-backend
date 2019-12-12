package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("nohp").
			Default("Unknown"),
		field.String("nama").
			Default("John Doe"),
		field.String("tgl_lahir").
			Default("00-00-00"),
		field.String("alamat").
			Default("Unknown"),
		field.Int("level").
			Positive(),
		field.String("image").
			Default(""),
		field.String("kategori_level").
			Default("pemula"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
