// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/confus1on/mbeledos/ent/bengkel"
	"github.com/facebookincubator/ent/dialect/sql"
)

// BengkelCreate is the builder for creating a Bengkel entity.
type BengkelCreate struct {
	config
	kode_bengkel   *string
	nama_bengkel   *string
	alamat_bengkel *string
	latitude       *float64
	longitude      *float64
	phonenumber    *string
}

// SetKodeBengkel sets the kode_bengkel field.
func (bc *BengkelCreate) SetKodeBengkel(s string) *BengkelCreate {
	bc.kode_bengkel = &s
	return bc
}

// SetNamaBengkel sets the nama_bengkel field.
func (bc *BengkelCreate) SetNamaBengkel(s string) *BengkelCreate {
	bc.nama_bengkel = &s
	return bc
}

// SetNillableNamaBengkel sets the nama_bengkel field if the given value is not nil.
func (bc *BengkelCreate) SetNillableNamaBengkel(s *string) *BengkelCreate {
	if s != nil {
		bc.SetNamaBengkel(*s)
	}
	return bc
}

// SetAlamatBengkel sets the alamat_bengkel field.
func (bc *BengkelCreate) SetAlamatBengkel(s string) *BengkelCreate {
	bc.alamat_bengkel = &s
	return bc
}

// SetLatitude sets the latitude field.
func (bc *BengkelCreate) SetLatitude(f float64) *BengkelCreate {
	bc.latitude = &f
	return bc
}

// SetLongitude sets the longitude field.
func (bc *BengkelCreate) SetLongitude(f float64) *BengkelCreate {
	bc.longitude = &f
	return bc
}

// SetPhonenumber sets the phonenumber field.
func (bc *BengkelCreate) SetPhonenumber(s string) *BengkelCreate {
	bc.phonenumber = &s
	return bc
}

// Save creates the Bengkel in the database.
func (bc *BengkelCreate) Save(ctx context.Context) (*Bengkel, error) {
	if bc.kode_bengkel == nil {
		return nil, errors.New("ent: missing required field \"kode_bengkel\"")
	}
	if bc.nama_bengkel == nil {
		v := bengkel.DefaultNamaBengkel
		bc.nama_bengkel = &v
	}
	if err := bengkel.NamaBengkelValidator(*bc.nama_bengkel); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"nama_bengkel\": %v", err)
	}
	if bc.alamat_bengkel == nil {
		return nil, errors.New("ent: missing required field \"alamat_bengkel\"")
	}
	if err := bengkel.AlamatBengkelValidator(*bc.alamat_bengkel); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"alamat_bengkel\": %v", err)
	}
	if bc.latitude == nil {
		return nil, errors.New("ent: missing required field \"latitude\"")
	}
	if bc.longitude == nil {
		return nil, errors.New("ent: missing required field \"longitude\"")
	}
	if bc.phonenumber == nil {
		return nil, errors.New("ent: missing required field \"phonenumber\"")
	}
	return bc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BengkelCreate) SaveX(ctx context.Context) *Bengkel {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bc *BengkelCreate) sqlSave(ctx context.Context) (*Bengkel, error) {
	var (
		builder = sql.Dialect(bc.driver.Dialect())
		b       = &Bengkel{config: bc.config}
	)
	tx, err := bc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	insert := builder.Insert(bengkel.Table).Default()
	if value := bc.kode_bengkel; value != nil {
		insert.Set(bengkel.FieldKodeBengkel, *value)
		b.KodeBengkel = *value
	}
	if value := bc.nama_bengkel; value != nil {
		insert.Set(bengkel.FieldNamaBengkel, *value)
		b.NamaBengkel = *value
	}
	if value := bc.alamat_bengkel; value != nil {
		insert.Set(bengkel.FieldAlamatBengkel, *value)
		b.AlamatBengkel = *value
	}
	if value := bc.latitude; value != nil {
		insert.Set(bengkel.FieldLatitude, *value)
		b.Latitude = *value
	}
	if value := bc.longitude; value != nil {
		insert.Set(bengkel.FieldLongitude, *value)
		b.Longitude = *value
	}
	if value := bc.phonenumber; value != nil {
		insert.Set(bengkel.FieldPhonenumber, *value)
		b.Phonenumber = *value
	}

	id, err := insertLastID(ctx, tx, insert.Returning(bengkel.FieldID))
	if err != nil {
		return nil, rollback(tx, err)
	}
	b.ID = int(id)
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return b, nil
}
