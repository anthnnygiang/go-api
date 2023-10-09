package web

import (
	"anthnnygiang/api-template/app"
	"encoding/json"
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
	switch r.Method {
	case http.MethodPost:
		//retrieve the request body
		payload := struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}{}
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			http.Error(w, "error: decode", http.StatusBadRequest)
		}
		defer r.Body.Close()

		//create the user
		id := uuid.New()
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 12)
		if err != nil {
			fmt.Printf("%v", err)
		}
		userData := app.User{
			ID:           id,
			CreatedAt:    time.Now(),
			Email:        payload.Email,
			PasswordHash: passwordHash,
			Activated:    false,
		}
		newUser, err := s.UserService.CreateUser(&userData)
		if err != nil {
			fmt.Printf("%v", err)
		}
		token := app.Token{
			TokenHash: [32]byte{},
			Scope:     app.ScopeActivate,
			Expiry:    time.Time{},
		}
		//send the activation email
		email := app.ActivationEmail{To: newUser.Email, Token: token}
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

func (s *Server) HandleActivateUser(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/activate" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodPut:
		//Handle the PUT request...

	case http.MethodOptions:
		w.Header().Set("Allow", "PUT, OPTIONS")
		w.WriteHeader(http.StatusNoContent)

	default:
		w.Header().Set("Allow", "PUT, OPTIONS")
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
