package controllers

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/labstack/echo/v4"
	"github.com/mahmudinashar/go/graph"
	"github.com/mahmudinashar/go/graph/generated"
)

func (server *Server) Graphql(c echo.Context) error {
	graphqlHandler := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graph.Resolver{DB: server.DB}},
		),
	)

	graphqlHandler.ServeHTTP(c.Response(), c.Request())
	return nil
}
