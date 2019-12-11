package repository

import (
	"context"
	"fmt"

	"github.com/confus1on/mbeledos/ent"
	"github.com/confus1on/mbeledos/ent/user"
	_ "github.com/lib/pq"

	"github.com/confus1on/mbeledos/userservice"
)

type PostgreSQLUserRepository struct {
	DB *ent.Client
}

func NewPostgreSQLUserRepository(db *ent.Client) userservice.Repository {
	return &PostgreSQLUserRepository{DB: db}
}

func (ur *PostgreSQLUserRepository) Login(ctx context.Context, phonenumber string) (res *ent.User, err error) {
	res, err = ur.DB.User.
		Query().
		Where(user.NohpEQ(phonenumber)).
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("Failed Querying user : %v", err.Error())
	}

	return res, nil
}

func (ur *PostgreSQLUserRepository) Register(ctx context.Context, user *ent.User) error {
	_, err := ur.DB.User.
		Create().
		SetNohp(user.Nohp).
		SetNama(user.Nama).
		SetTglLahir(user.TglLahir).
		SetAlamat(user.Alamat).
		Save(ctx)

	if err != nil {
		return err
	}

	return nil
}
