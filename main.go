package main

import (
	"anthnnygiang/api-template/postgres"
	"anthnnygiang/api-template/postmark"
	"anthnnygiang/api-template/web"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
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

	// Create an HTTP client per service where needed
	HTTPClient := &http.Client{}

	//Initialize services
	userService := &postgres.UserService{
		DB: db,
	}
	emailService := &postmark.EmailService{
		HTTPClient: HTTPClient,
		APIKey:     os.Getenv("POSTMARK_SERVER_TOKEN"),
	}

	/* Initialize server
	Services must implement all methods. This allows for
	different implementations to be easily interchangeable.
	*/
	server := &web.Server{
		UserService:  userService,
		EmailService: emailService,
	}

	//move to chi
	mux := http.NewServeMux()
	mux.HandleFunc("/", server.HandleIndex)
	mux.HandleFunc("/signup", server.HandleSignUp)
	mux.HandleFunc("/activate", server.HandleActivateUser)
	mux.HandleFunc("/signin", server.HandleSignIn)

	const PORT = "4000"
	fmt.Printf("server @ localhost:%s\n", PORT)
	err = http.ListenAndServe(":"+PORT, mux)
	log.Fatal(err)
}
