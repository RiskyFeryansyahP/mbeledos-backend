package user

import (
	"github.com/confus1on/mbeledos/models"
)

type Repository interface {
	Login(phonenumber string) (*models.User, error)
	Register(user *models.User) error
}
