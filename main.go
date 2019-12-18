package main

import (
	"fmt"
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/gomodule/redigo/redis"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/valyala/fasthttp"

	"github.com/confus1on/mbeledos/config"
	"github.com/confus1on/mbeledos/ent"

	// route package user
	handleruser "github.com/confus1on/mbeledos/userservice/handler"
	repouser "github.com/confus1on/mbeledos/userservice/repository"
	usecaseuser "github.com/confus1on/mbeledos/userservice/usecase"

	// route package bengkel
	handlerbengkel "github.com/confus1on/mbeledos/bengkelservice/handler"
	repobengkel "github.com/confus1on/mbeledos/bengkelservice/repository"
	usecasebengkel "github.com/confus1on/mbeledos/bengkelservice/usecase"
)

func Hello(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Hello World")
}

var DB *ent.Client
var ConnRedis redis.Conn

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("did not find the file .env")
	}

	// initialize router using fasthttprouter
	router := fasthttprouter.New()

	// call pool of redis
	pool := config.NewPool()
	ConnRedis = pool.Get()
	defer ConnRedis.Close()

	// call new config database
	DB = config.New()
	defer DB.Close()

	// initalize route
	// userservice route
	postgresqlUserRepository := repouser.NewPostgreSQLUserRepository(DB)
	userUsecase := usecaseuser.NewUserUsercase(postgresqlUserRepository)
	user := handleruser.NewUserHandler(userUsecase)

	// bengkelservice route
	postgresqlBengkelRepository := repobengkel.NewPostgreSQLBengkelRepository(DB)
	bengkelUsecase := usecasebengkel.NewBengkelUsecase(postgresqlBengkelRepository)
	bengkelHandler := handlerbengkel.NewBengkelHandler(bengkelUsecase)

	router.POST("/user/login", user.Login)
	router.POST("/user/register", user.Register)
	router.POST("/user/verification", user.SendVerificationCode)

	router.GET("/bengkel/all", bengkelHandler.GetAllDataBengkel)
	router.GET("/bengkel/kode/:kode_bengkel", bengkelHandler.GetDataBengkel)

	log.Println("Server Running on http://127.0.0.1:8080")
	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}
