package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo/v4"
	"github.com/suradidchao/listenfield/handler"
	"github.com/suradidchao/listenfield/repo/farm"
	"github.com/suradidchao/listenfield/usecase"
)

func main() {

	const (
		mysqlUser   = "zocialeye"
		mysqlPass   = "zocialeye"
		mysqlHost   = "localhost"
		mysqlPort   = 3306
		mysqlDBName = "lb"
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
	e := echo.New()

	farmGroup := e.Group("/farms")
	farmGroup.POST("/", farmHandler.CreateFarm)
	e.Logger.Fatal(e.Start(":8000"))

}
