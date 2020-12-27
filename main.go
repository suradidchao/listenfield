package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/suradidchao/listenfield/handler"
	"github.com/suradidchao/listenfield/internal/jwtgen"
	"github.com/suradidchao/listenfield/repo/farm"
	"github.com/suradidchao/listenfield/repo/farmworker"
	"github.com/suradidchao/listenfield/repo/user"
	"github.com/suradidchao/listenfield/usecase"
)

func main() {

	env := "local"
	configPath := fmt.Sprintf("./config/%s/", env)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)
	fmt.Println(configPath)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	var (
		mysqlUser   = viper.GetString("mysql.user")
		mysqlPass   = viper.GetString("mysql.pass")
		mysqlHost   = viper.GetString("mysql.host")
		mysqlPort   = viper.GetInt("mysql.port")
		mysqlDBName = viper.GetString("mysql.db")
	)

	apiSecret := viper.GetString("api.secret")

	fmt.Println("MysqlHost:", mysqlHost)
	fmt.Println("MysqlPort:", mysqlPort)
	fmt.Println("MysqlUser:", mysqlUser)
	fmt.Println("MysqlDBName:", mysqlDBName)

	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", mysqlUser, mysqlPass, mysqlHost, mysqlPort, mysqlDBName)
	mysqlDB, err := sql.Open("mysql", mysqlURI+"?parseTime=true&charset=utf8")
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to mysql database with err: %s", err))
	}
	defer mysqlDB.Close()

	farmWorkerSQLAdapter := farmworker.NewMySQLAdapter(mysqlDB)
	farmWorkerRepo := farmworker.NewRepo(farmWorkerSQLAdapter)

	farmSQLAdapter := farm.NewMySQLAdapter(mysqlDB)
	farmRepo := farm.NewRepo(farmSQLAdapter)
	farmUsecase := usecase.NewFarmUsecase(farmRepo, farmWorkerRepo)
	farmHandler := handler.NewFarmHandler(farmUsecase)

	userSQLAdapter := user.NewMySQLAdapter(mysqlDB)
	userRepo := user.NewRepo(userSQLAdapter, farmSQLAdapter, farmWorkerSQLAdapter)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	jwtGenerator := jwtgen.NewJWTGenerator(apiSecret)
	authUsecase := usecase.NewAuthUsecase(userRepo, jwtGenerator)
	authHandler := handler.NewAuthHandler(authUsecase)
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	e.Use(middleware.Recover())

	isLogin := middleware.JWT([]byte(apiSecret))

	farmGroup := e.Group("/farms")
	farmGroup.POST("", farmHandler.CreateFarm)
	farmGroup.POST("/:farm_id/workers", farmHandler.AddWorker, isLogin)
	farmGroup.DELETE("/:farm_id/workers/:farmworker_id", farmHandler.DeleteWorker, isLogin)
	farmGroup.GET("/:farm_id/workers", farmHandler.GetAllWorkers, isLogin)

	e.POST("/authenticate", authHandler.Authenticate)
	e.POST("/users", userHandler.Create)

	e.Logger.Fatal(e.Start(":8000"))

}
