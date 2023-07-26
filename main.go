package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/giovanisilqueirasantos/coinconv/config"
	"github.com/giovanisilqueirasantos/coinconv/handler"
	"github.com/giovanisilqueirasantos/coinconv/repo"
	"github.com/giovanisilqueirasantos/coinconv/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	conf, err := config.GetConf("./config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	dbConn, err := sql.Open(`mysql`, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.Database.User, conf.Database.Pass, conf.Database.Host, conf.Database.Port, conf.Database.Name))

	if err != nil {
		log.Fatal(err)
	}

	err = dbConn.Ping()

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()

	e.Use(middleware.CORS())

	repo := repo.NewMysqlRepo(dbConn)
	service := service.NewService(repo)

	handler.NewHttpHandler(e, service)

	log.Fatal(e.Start(conf.Server.Address))
}
