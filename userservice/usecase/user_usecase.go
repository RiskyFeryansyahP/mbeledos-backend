package usecase

import (
	"context"

	"github.com/confus1on/mbeledos/ent"
	"github.com/confus1on/mbeledos/userservice"
)

type UserUsecase struct {
	UserRepository userservice.Repository
}

func NewUserUsercase(ur userservice.Repository) userservice.Usecase {
	return &UserUsecase{UserRepository: ur}
}

func (uc *UserUsecase) Login(ctx context.Context, phonenumber string) (res *ent.User, err error) {
	res, err = uc.UserRepository.Login(ctx, phonenumber)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (uc *UserUsecase) Register(ctx context.Context, user *ent.User) error {
	err := uc.UserRepository.Register(ctx, user)
	if err != nil {
		return err
	}
	return nil
}