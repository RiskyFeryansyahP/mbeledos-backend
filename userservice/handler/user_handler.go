package handler

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/confus1on/mbeledos/ent"
	"github.com/confus1on/mbeledos/userservice"
	"github.com/valyala/fasthttp"
)

type UserHandler struct {
	UserUsecase userservice.Usecase
}

type Message struct {
	Status int `json:"status"`
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

func (uh *UserHandler) SendVerificationCode(ctx *fasthttp.RequestCtx) {
	var user *ent.User

	json.Unmarshal(ctx.Request.Body(), &user)

	otp, err := uh.UserUsecase.SendOTPVerification(ctx, user.Nohp)
	if err != nil {
		log.Println(err.Error())
		return
	}

	json.NewEncoder(ctx).Encode(map[string]string{
		"Message": "Success Send Verification Code",
		"OTP":     strconv.Itoa(int(otp)),
	})
}

func (uh *UserHandler) UpdateUser(ctx *fasthttp.RequestCtx) {
	var user ent.User

	body := ctx.Request.Body()

	json.Unmarshal(body, &user)

	err := uh.UserUsecase.UpdateUser(ctx, user)
	if err != nil {
		log.Println(err)
	}

	json.NewEncoder(ctx).Encode(&Message{Status: 200})
}
