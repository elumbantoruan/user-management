package main

import (
	"edison-takehome/pkg/handler"
	"edison-takehome/pkg/repository/postgres"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	dbserver := os.Getenv("DBSERVER")

	pg, err := postgres.NewPostgresDB(dbserver, 5432, "postgres", "postgres")
	if err != nil {
		log.Fatal(err)
	}
	defer pg.Close()

	signUpHandler := handler.NewSignup(pg)
	loginHandler := handler.NewLogin(pg)
	userHandler := handler.NewUserHandler(pg)

	gin.SetMode(gin.ReleaseMode)
	route := gin.New()

	route.POST("/usermanagement/v1/signup", signUpHandler.UserSignup)
	route.POST("/usermanagement/v1/login", loginHandler.UserLogin)
	route.GET("/usermanagement/v1/users", userHandler.GetUserList)
	route.PUT("/usermanagement/v1/users", userHandler.UpdateUser)

	err = route.Run(":8088")
	if err != nil {
		log.Fatal(err)
	}
}
