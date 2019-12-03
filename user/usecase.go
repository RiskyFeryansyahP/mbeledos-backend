package user

import (
	"github.com/confus1on/mbeledos/models"
)

type Usecase interface {
	Login(phonenumber string) (*models.User, error)
	Register(user *models.User) error
}
