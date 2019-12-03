package main

import (
	"fmt"
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/confus1on/mbeledos/config"

	// route package
	"github.com/confus1on/mbeledos/user/handler"
	"github.com/confus1on/mbeledos/user/repository"
	"github.com/confus1on/mbeledos/user/usecase"
)

func Hello(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Hello World")
}

var DB *mongo.Database

func main() {

	// initialize router using fasthttprouter
	router := fasthttprouter.New()

	// call new config database
	DB := config.New()

	// initalize route
	mongoUserRepository := repository.NewMongoUserRepository(DB)
	userUsecase := usecase.NewUserUsercase(mongoUserRepository)
	user := handler.NewUserHandler(userUsecase)

	router.POST("/user/login", user.Login)
	router.POST("/user/register", user.Register)

	log.Println("Server Running on http://127.0.0.1:8080")
	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}
