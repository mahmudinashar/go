package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type Server struct {
	DB     *gorm.DB
	Router *echo.Echo
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error
	err = godotenv.Load()

	if Dbdriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

		// connect to database
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("Connected to %s database", Dbdriver)
		}
	} else {
		fmt.Println("System only allow MySQL database driver, contant mahmudinashar@yahoo.co.id")
	}

	// database migration (just for) make sure that table is exist!
	// server.DB.Debug().AutoMigrate(&models.User{}, &models.Post{})
	// server.Router = mux.NewRouter()

	server.Router = echo.New()
	server.initializeRoutes()
}

// use gorilla/mux
func (server *Server) Run(addr string) {
	fmt.Println(", application running in ", os.Getenv("SYS_PORT"))
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

// use echo
func (server *Server) Start(addr string) error {
	return server.Router.Start(":" + addr)
}
