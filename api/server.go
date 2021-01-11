package api

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mahmudinashar/go/api/controllers"
	"github.com/mahmudinashar/go/api/routes"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("YES!, .env file found")
	}
}

func Run() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("OMG!, Error on .env, %v", err)
	} else {
		fmt.Println("GOT .env values")
	}

	app := controllers.Initialize()
	routes.Routing(app)

	err = app.Start(os.Getenv("SYS_PORT"))

	if err != nil {
		log.Fatalf("OMG!, Error : %v", err)
	}

}
