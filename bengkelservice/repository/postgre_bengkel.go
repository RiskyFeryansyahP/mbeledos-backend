package repository

import (
	"context"
	"fmt"

	"github.com/confus1on/mbeledos/bengkelservice"

	"github.com/confus1on/mbeledos/ent"
	"github.com/confus1on/mbeledos/ent/bengkel"
)

type PostgreSQLBengkelRepository struct {
	DB *ent.Client
}

func NewPostgreSQLBengkelRepository(DB *ent.Client) bengkelservice.RepositoryBengkel {
	return &PostgreSQLBengkelRepository{DB: DB}
}

func (br *PostgreSQLBengkelRepository) ShowAllBengkel(ctx context.Context) ([]*ent.Bengkel, error) {
	result, err := br.DB.Bengkel.
		Query().
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("Error when get all data bengkel %v", err.Error())
	}

	return result, nil
}

func (br *PostgreSQLBengkelRepository) GetSpecificationBengkel(ctx context.Context, kode_bengkel string) (*ent.Bengkel, error) {
	result, err := br.DB.Bengkel.
		Query().
		Where(bengkel.KodeBengkelEQ(kode_bengkel)).
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("Error when get specification bengkel : %v", err.Error())
	}

	return result, nil
}

func (br *PostgreSQLBengkelRepository) SigninBengkelUser(ctx context.Context, phone string) (*ent.Bengkel, error) {
	result, err := br.DB.Bengkel.
		Query().
		Where(bengkel.PhonenumberEQ(phone)).
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("Error when login bengkel %v ", err.Error())
	}

	return result, nil
}
