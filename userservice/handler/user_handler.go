package handler

import (
	"encoding/json"
	"log"

	"github.com/confus1on/mbeledos/ent"
	"github.com/confus1on/mbeledos/userservice"
	"github.com/valyala/fasthttp"
)

type UserHandler struct {
	UserUsecase userservice.Usecase
}

func NewUserHandler(uc userservice.Usecase) *UserHandler {
	return &UserHandler{UserUsecase: uc}
}

func (uh *UserHandler) Login(ctx *fasthttp.RequestCtx) {
	var user ent.User

	err := json.Unmarshal(ctx.Request.Body(), &user)
	if err != nil {
		log.Println("user handler : error when unmarshal", err.Error())
	}

	res, err := uh.UserUsecase.Login(ctx, user.Nohp)
	if err != nil {
		log.Println("mongo_user : Can't find data in users with phone number : "+user.Nohp, err.Error())
	}

	json.NewEncoder(ctx).Encode(res)
}

func (uh *UserHandler) Register(ctx *fasthttp.RequestCtx) {
	user := &ent.User{}

	json.Unmarshal(ctx.Request.Body(), &user)

	err := uh.UserUsecase.Register(ctx, user)
	if err != nil {
		log.Println("mongo_user : something wrong went inserted", err.Error())
		return
	}

	json.NewEncoder(ctx).Encode(user)
}
