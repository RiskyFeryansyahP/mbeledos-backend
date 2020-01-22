package userservice

import (
	"context"

	"github.com/confus1on/mbeledos/ent"
)

type Repository interface {
	Login(ctx context.Context, phonenumber string) (*ent.User, error)
	Register(ctx context.Context, user *ent.User) error
	SendVerification(ctx context.Context, OTP int32, phonenumber string) error
	UpdateUser(ctx context.Context, profile ent.User) error
}
