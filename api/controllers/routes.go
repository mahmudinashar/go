package controllers

import (
	"os"

	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) initializeRoutes() {
	// s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")
	// s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")
	// s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	// s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	// s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	// s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	// s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")
	// s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.CreatePost)).Methods("POST")
	// s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.GetPosts)).Methods("GET")
	// s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(s.GetPost)).Methods("GET")
	// s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdatePost))).Methods("PUT")
	// s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareAuthentication(s.DeletePost)).Methods("DELETE")

	e := s.Router.Group("")
	e.Use(middleware.Gzip())

	e.GET("/", s.Home)
	e.POST("/login", s.Login)
	e.GET("/users", s.GetUsers)
	e.GET("/hello", s.Home)

	// routing group authJWT use after user login
	authJWT := s.Router.Group("")
	authJWT.Use(middleware.Gzip())
	authJWT.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: os.Getenv("API_SECRET_METHOD"),
		SigningKey:    []byte(os.Getenv("API_SECRET")),
	}))

	authJWT.PUT("/users", s.UpdateUser)

}
