package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/confus1on/mbeledos/ent"
	"github.com/confus1on/mbeledos/ent/user"
	_ "github.com/lib/pq"

	"github.com/confus1on/mbeledos/models"
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
		SetLevel(user.Level).
		SetImage(user.Image).
		SetKategoriLevel(user.KategoriLevel).
		Save(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (PostgreSQLUserRepository) SendVerification(ctx context.Context, OTP int32, phonenumber string) error {
	nexmoSMS := models.RequestNexmoSMS{
		API_KEY:    os.Getenv("API_KEY"),
		API_SECRET: os.Getenv("API_SECRET"),
		To:         phonenumber,
		From:       "Mbeledos",
		Text:       "Kode Verifikasi Anda : " + strconv.Itoa(int(OTP)),
	}

	b, err := json.Marshal(nexmoSMS)
	if err != nil {
		return fmt.Errorf("something went wrong in verification : %v", err.Error())
	}

	_, err = http.Post("https://rest.nexmo.com/sms/json", "application/json", bytes.NewBuffer(b))
	if err != nil {
		return fmt.Errorf("Error when POST in nexmo : %v", err.Error())
	}

	return nil
}
