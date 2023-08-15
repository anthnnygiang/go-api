package main

import (
	"9z/go-api-template/postgres"
	"9z/go-api-template/web"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	//load environment variables
	//TODO: use setenv
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

	/** Initialize server
	It is not important how services are implemented, as long
	as they implement all methods. This allows for different
	service implementations to be easily interchangeable.
	*/
	server := &web.Server{
		UserService:    userService,
		SessionService: sessionService,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", server.HandleIndex)
	mux.HandleFunc("/user", server.HandleUser)

	const PORT = "4000"
	fmt.Printf("server @ localhost:%s\n", PORT)
	err = http.ListenAndServe(":"+PORT, mux)
	log.Fatal(err)
}
