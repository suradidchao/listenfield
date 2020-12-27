package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/suradidchao/listenfield/handler"
	"github.com/suradidchao/listenfield/repo/farm"
	"github.com/suradidchao/listenfield/repo/user"
	"github.com/suradidchao/listenfield/usecase"
)

func main() {

	const (
		mysqlUser   = "listenfield"
		mysqlPass   = "listenfield"
		mysqlHost   = "localhost"
		mysqlPort   = 3306
		mysqlDBName = "listenfield"
	)

	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", mysqlUser, mysqlPass, mysqlHost, mysqlPort, mysqlDBName)
	mysqlDB, err := sql.Open("mysql", mysqlURI+"?parseTime=true&charset=utf8")
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to mysql database with err: %s", err))
	}
	defer mysqlDB.Close()

	farmSQLAdapter := farm.NewMySQLAdapter(mysqlDB)
	farmRepo := farm.NewRepo(farmSQLAdapter)
	farmUsecase := usecase.NewFarmUsecase(farmRepo)
	farmHandler := handler.NewFarmHandler(farmUsecase)

	userSQLAdapter := user.NewMySQLAdapter(mysqlDB)
	userRepo := user.NewRepo(userSQLAdapter)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	authorizeHandler := handler.NewAuthorizeHandler(userUsecase)
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	e.Use(middleware.Recover())

	farmGroup := e.Group("/farms")
	farmGroup.POST("", farmHandler.CreateFarm)

	e.POST("/authorize", authorizeHandler.Authorize)
	e.POST("/users", userHandler.Create)

	e.Logger.Fatal(e.Start(":8000"))

}
