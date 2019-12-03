package handler

import (
	"encoding/json"
	"log"

	"github.com/confus1on/mbeledos/models"
	"github.com/confus1on/mbeledos/user"
	"github.com/valyala/fasthttp"
)

type UserHandler struct {
	UserUsecase user.Usecase
}

func NewUserHandler(uc user.Usecase) *UserHandler {
	return &UserHandler{UserUsecase: uc}
}

func (uh *UserHandler) Login(ctx *fasthttp.RequestCtx) {
	var user models.User

	err := json.Unmarshal(ctx.Request.Body(), &user)
	if err != nil {
		log.Println("user handler : error when unmarshal", err.Error())
	}

	res, err := uh.UserUsecase.Login(user.PhoneNumber)
	if err != nil {
		log.Println("mongo_user : Can't find data in users with phone number : "+user.PhoneNumber, err.Error())
	}

	json.NewEncoder(ctx).Encode(res)
}

func (uh *UserHandler) Register(ctx *fasthttp.RequestCtx) {
	user := &models.User{}

	json.Unmarshal(ctx.Request.Body(), &user)

	err := uh.UserUsecase.Register(user)
	if err != nil {
		log.Println("mongo_user : something wrong went inserted", err.Error())
	}

	json.NewEncoder(ctx).Encode(user)
}
