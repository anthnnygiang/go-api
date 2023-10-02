package web

import (
	"anthnnygiang/api-template/app"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func (s *Server) HandleSignUp(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/signup" {
		http.NotFound(w, r)
		return
	}
	//Common code for all requests...

	switch r.Method {
	case http.MethodPost:
		/* Handle the POST request
		1. Create a new uuid
		2. Create the user struct
		3. Call the service to create the user
		4. Send activation email
		5. On success, return success response. Or on error, return error response
		*/

		//Construct, and create the user
		id := uuid.New()
		//Retrieve values from request
		passwordHash, err := bcrypt.GenerateFromPassword([]byte("a password"), 12)
		if err != nil {
			fmt.Printf("%v", err)
		}
		userData := app.User{
			ID:           id,
			CreatedAt:    time.Now(), //Create time here or in db?
			Email:        "example@example.com",
			PasswordHash: passwordHash,
			Activated:    false,
		}
		newUser, err := s.UserService.CreateUser(&userData)
		if err != nil {
			fmt.Printf("%v", err)
		}
		//Construct, and send the activation email
		email := app.ActivationEmail{To: newUser.Email}
		_, err = s.EmailService.SendActivationEmail(email)
		if err != nil {
			fmt.Printf("%v", err)
		}

		//Write JSON response
		fmt.Fprintf(w, "%v", newUser.Email)

	case http.MethodOptions:
		w.Header().Set("Allow", "POST, OPTIONS")
		w.WriteHeader(http.StatusNoContent)

	default:
		w.Header().Set("Allow", "POST, OPTIONS")
		http.Error(w, "Error: method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Server) HandleSignIn(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/signin" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodPost:
		//Handle the POST request...
		//Create a new session

	case http.MethodOptions:
		w.Header().Set("Allow", "POST, OPTIONS")
		w.WriteHeader(http.StatusNoContent)

	default:
		w.Header().Set("Allow", "POST, OPTIONS")
		http.Error(w, "Error: method not allowed", http.StatusMethodNotAllowed)
	}
}
