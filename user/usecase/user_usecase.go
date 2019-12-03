package usecase

import (
	"github.com/confus1on/mbeledos/models"
	"github.com/confus1on/mbeledos/user"
)

type UserUsecase struct {
	UserRepository user.Repository
}

func NewUserUsercase(ur user.Repository) user.Usecase {
	return &UserUsecase{UserRepository: ur}
}

func (uc *UserUsecase) Login(phonenumber string) (res *models.User, err error) {
	res, err = uc.UserRepository.Login(phonenumber)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (uc *UserUsecase) Register(user *models.User) error {
	err := uc.UserRepository.Register(user)
	if err != nil {
		return err
	}
	return nil
}
