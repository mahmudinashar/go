package api

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mahmudinashar/go/api/controllers"
)

// push controllers/base.go > 'Server struct', to variable 'server'
var server = controllers.Server{}

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

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	// Activate this for first time running, then setting seeder in github.com/mahmudinashar/go/api/seed
	// seed.Load(server.DB)

	// use gorilla/mux
	// server.Run(os.Getenv("SYS_PORT"))

	// use echo
	err = server.Start(os.Getenv("SYS_PORT"))

	if err != nil {
		log.Fatalf("OMG!, Error : %v", err)
	}

}
