package controllers

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	"github.com/mahmudinashar/go/api/seed"
)

type Server struct {
	Router *echo.Echo
	DB     *gorm.DB
}

func Initialize() *Server {
	var database *gorm.DB
	if os.Getenv("DB_DRIVER") == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

		db, err := gorm.Open(os.Getenv("DB_DRIVER"), DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", os.Getenv("DB_DRIVER"))
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("Connected to %s database", os.Getenv("DB_DRIVER"))
		}

		database = db

		//  +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
		//   comment `seed.Load` bellow, when you are running this app for the second time or more.
		//   this script is use for import database schema+data into target database (MySQL,
		//   Postgres, SQLLite) or other database that support by GORM (gorm.io/gorm)
		//  +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

		seed.Load(db)

	} else {
		fmt.Println("System only allow MySQL database driver")
	}

	return &Server{
		Router: echo.New(),
		DB:     database,
	}
}

func (server *Server) Start(addr string) error {
	return server.Router.Start(":" + addr)
}
