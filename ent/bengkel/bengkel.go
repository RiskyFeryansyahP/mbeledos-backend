// Code generated by entc, DO NOT EDIT.

package bengkel

import (
	"github.com/confus1on/mbeledos/ent/schema"
)

const (
	// Label holds the string label denoting the bengkel type in the database.
	Label = "bengkel"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldKodeBengkel holds the string denoting the kode_bengkel vertex property in the database.
	FieldKodeBengkel = "kode_bengkel"
	// FieldNamaBengkel holds the string denoting the nama_bengkel vertex property in the database.
	FieldNamaBengkel = "nama_bengkel"
	// FieldAlamatBengkel holds the string denoting the alamat_bengkel vertex property in the database.
	FieldAlamatBengkel = "alamat_bengkel"
	// FieldLatitude holds the string denoting the latitude vertex property in the database.
	FieldLatitude = "latitude"
	// FieldLongitude holds the string denoting the longitude vertex property in the database.
	FieldLongitude = "longitude"
	// FieldPhonenumber holds the string denoting the phonenumber vertex property in the database.
	FieldPhonenumber = "phonenumber"

	// Table holds the table name of the bengkel in the database.
	Table = "bengkels"
)

// Columns holds all SQL columns are bengkel fields.
var Columns = []string{
	FieldID,
	FieldKodeBengkel,
	FieldNamaBengkel,
	FieldAlamatBengkel,
	FieldLatitude,
	FieldLongitude,
	FieldPhonenumber,
}

var (
	fields = schema.Bengkel{}.Fields()

	// descNamaBengkel is the schema descriptor for nama_bengkel field.
	descNamaBengkel = fields[1].Descriptor()
	// DefaultNamaBengkel holds the default value on creation for the nama_bengkel field.
	DefaultNamaBengkel = descNamaBengkel.Default.(string)
	// NamaBengkelValidator is a validator for the "nama_bengkel" field. It is called by the builders before save.
	NamaBengkelValidator = descNamaBengkel.Validators[0].(func(string) error)

	// descAlamatBengkel is the schema descriptor for alamat_bengkel field.
	descAlamatBengkel = fields[2].Descriptor()
	// AlamatBengkelValidator is a validator for the "alamat_bengkel" field. It is called by the builders before save.
	AlamatBengkelValidator = descAlamatBengkel.Validators[0].(func(string) error)
)
