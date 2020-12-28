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
	customMiddleware "github.com/suradidchao/listenfield/middleware"
	"github.com/suradidchao/listenfield/repo/activity"
	"github.com/suradidchao/listenfield/repo/farm"
	"github.com/suradidchao/listenfield/repo/farmworker"
	"github.com/suradidchao/listenfield/repo/field"
	"github.com/suradidchao/listenfield/repo/tractor"
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

	tractorSQLAdapter := tractor.NewMySQLAdapter(mysqlDB)
	tractorRepo := tractor.NewRepo(tractorSQLAdapter)

	farmWorkerSQLAdapter := farmworker.NewMySQLAdapter(mysqlDB)
	farmWorkerRepo := farmworker.NewRepo(farmWorkerSQLAdapter)

	fieldSQLAdapter := field.NewMySQLAdapter(mysqlDB)
	fieldRepo := field.NewRepo(fieldSQLAdapter)

	activitySQLAdapter := activity.NewMySQLAdapter(mysqlDB)
	activityRepo := activity.NewRepo(activitySQLAdapter)

	farmSQLAdapter := farm.NewMySQLAdapter(mysqlDB)
	farmRepo := farm.NewRepo(farmSQLAdapter)
	farmUsecase := usecase.NewFarmUsecase(farmRepo, farmWorkerRepo, tractorRepo, fieldRepo, activityRepo)
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

	farmGroup := e.Group("/farms", isLogin)
	farmGroup.POST("", farmHandler.CreateFarm)
	farmGroup.POST("/:farm_id/workers", farmHandler.AddWorker, customMiddleware.AuthorizeFarmOwnerAccess)
	farmGroup.DELETE("/:farm_id/workers/:farmworker_id", farmHandler.DeleteWorker, customMiddleware.AuthorizeFarmOwnerAccess)
	farmGroup.GET("/:farm_id/workers", farmHandler.GetAllWorkers, customMiddleware.AuthorizeFarmOwnerAccess)
	farmGroup.POST("/:farm_id/tractors", farmHandler.AddTractor, customMiddleware.AuthorizeFarmOwnerAccess)
	farmGroup.DELETE("/:farm_id/tractors/:tractor_id", farmHandler.DeleteTractor, customMiddleware.AuthorizeFarmOwnerAccess)
	farmGroup.PUT("/:farm_id/tractors/:tractor_id", farmHandler.UpdateTractor, customMiddleware.AuthorizeFarmOwnerAccess)
	farmGroup.POST("/:farm_id/fields", farmHandler.AddField, customMiddleware.AuthorizeFarmOwnerAccess)
	farmGroup.DELETE("/:farm_id/fields/:field_id", farmHandler.DeleteField, customMiddleware.AuthorizeFarmOwnerAccess)
	farmGroup.PUT("/:farm_id/fields/:field_id", farmHandler.UpdateField, customMiddleware.AuthorizeFarmOwnerAccess)
	farmGroup.POST("/:farm_id/fields/:field_id/activities", farmHandler.AddActivity, customMiddleware.AuthorizeFarmOwnerAndWorkerAccess)
	farmGroup.GET("/:farm_id/costsummary", farmHandler.GetCostSummary, customMiddleware.AuthorizeFarmOwnerAccess)

	e.POST("/authenticate", authHandler.Authenticate)
	e.POST("/users", userHandler.Create)

	e.Logger.Fatal(e.Start(":8000"))

}
