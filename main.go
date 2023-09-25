package main

import (
	"anthnnygiang/api-template/postgres"
	"anthnnygiang/api-template/web"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	//load environment variables
	//Move to use os.Setenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Open a connection to the database
	db, err := postgres.Open()
	if err != nil {
		log.Fatal("Error DB")
	}
	defer db.Close()

	//Initialize services
	userService := &postgres.UserService{
		DB: db,
	}
	sessionService := &postgres.SessionService{
		DB: db,
	}

	/* Initialize server
	Services must implement all methods. This allows for
	different implementations to be easily interchangeable.
	*/
	server := &web.Server{
		UserService:    userService,
		SessionService: sessionService,
	}

	//Move to gorilla/mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", server.HandleIndex)
	mux.HandleFunc("/signup", server.HandleSignUp)
	mux.HandleFunc("/signin", server.HandleSignIn)

	const PORT = "4000"
	fmt.Printf("server @ localhost:%s\n", PORT)
	err = http.ListenAndServe(":"+PORT, mux)
	log.Fatal(err)
}
