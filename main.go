package main

import (
	"fmt"
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"

	"github.com/confus1on/mbeledos/config"
	"github.com/confus1on/mbeledos/ent"
	_ "github.com/lib/pq"

	// route package
	"github.com/confus1on/mbeledos/userservice/handler"
	"github.com/confus1on/mbeledos/userservice/repository"
	"github.com/confus1on/mbeledos/userservice/usecase"
)

func Hello(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Hello World")
}

var DB *ent.Client

func main() {

	// initialize router using fasthttprouter
	router := fasthttprouter.New()

	// call new config database
	DB = config.New()
	defer DB.Close()

	// initalize route
	mongoUserRepository := repository.NewPostgreSQLUserRepository(DB)
	userUsecase := usecase.NewUserUsercase(mongoUserRepository)
	user := handler.NewUserHandler(userUsecase)

	router.POST("/user/login", user.Login)
	router.POST("/user/register", user.Register)

	log.Println("Server Running on http://127.0.0.1:8080")
	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}
