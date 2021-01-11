package routes

import (
	"os"

	"github.com/labstack/echo/v4/middleware"
	ctrl "github.com/mahmudinashar/go/api/controllers"
)

func Routing(s *ctrl.Server) {
	e := s.Router.Group("")
	e.Use(middleware.Gzip())

	e.GET("/hello", s.Home)

	e.POST("/login", s.Login)
	e.GET("/json", s.Json)
	e.POST("/json/findById", s.Find)

	//  ++++++++++++++++++++++++++++++++++++
	// 	GRAPHQL REQUEST HANDLERS
	//	++++++++++++++++++++++++++++++++++++

	e.POST("/graphql", s.Graphql)

	//  ++++++++++++++++++++++++++++++++++++
	// 	authJWT IS USED AFTER AUTHENTICATED
	//	++++++++++++++++++++++++++++++++++++

	authJWT := s.Router.Group("")
	authJWT.Use(middleware.Gzip())
	authJWT.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: os.Getenv("API_SECRET_METHOD"),
		SigningKey:    []byte(os.Getenv("API_SECRET")),
	}))

	e.POST("/users", s.GetUsers)
	authJWT.POST("/users/create", s.CreateUser)
	authJWT.POST("/users/get", s.GetUser)
	authJWT.PUT("/users", s.UpdateUser)
	authJWT.DELETE("/users", s.DeleteUser)
}
