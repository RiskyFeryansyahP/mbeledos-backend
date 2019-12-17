package userservice

import (
	"context"

	"github.com/confus1on/mbeledos/ent"
)

type Usecase interface {
	Login(ctx context.Context, phonenumber string) (*ent.User, error)
	Register(ctx context.Context, user *ent.User) error
	SendOTPVerification(ctx context.Context, phonenumber string) (int32, error)
}
