package web

import (
	"anthnnygiang/api-template/app"
	"fmt"
	"github.com/rs/xid"
	"net/http"
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
		*/
		id := xid.New().String()

		//TODO: Get details from the request
		userData := app.User{ID: id, Email: "example@example.com", PasswordHash: "password hash"}
		newUser, err := s.UserService.CreateUser(&userData)
		if err != nil {
			fmt.Printf("%v", err)
		}

		err = s.EmailService.SendActivationEmail(newUser.Email)

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

	case http.MethodOptions:
		w.Header().Set("Allow", "POST, OPTIONS")
		w.WriteHeader(http.StatusNoContent)

	default:
		w.Header().Set("Allow", "POST, OPTIONS")
		http.Error(w, "Error: method not allowed", http.StatusMethodNotAllowed)
	}
}
