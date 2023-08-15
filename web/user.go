package web

import (
	"9z/go-api-template/app"
	"fmt"
	"github.com/rs/xid"
	"net/http"
)

func (s *Server) HandleUser(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/user" {
		http.NotFound(w, r)
		return
	}
	fmt.Printf("%s %s\n", r.Method, r.URL.Path)

	// Common code for all requests...

	switch r.Method {
	case http.MethodPost:
		// Handle the POST request...
		fmt.Fprintf(w, "handling POST request\n")
		id := xid.New().String()
		//TODO: Get details from the request
		userData := app.User{ID: id, Email: "example@example.com", PasswordHash: "password hash"}
		newUser, err := s.UserService.CreateUser(&userData)
		if err != nil {
			fmt.Printf("%v", err)
		}
		//Create a session
		newSession, rawToken, err := s.SessionService.CreateSession(newUser.ID)
		fmt.Printf("session: %v,\nraw token %v,\n", newSession, rawToken)

		//	Write JSON response
		fmt.Fprintf(w, "%v", newUser.Email)

	case http.MethodOptions:
		w.Header().Set("Allow", "POST, OPTIONS")
		w.WriteHeader(http.StatusNoContent)

	default:
		w.Header().Set("Allow", "POST, OPTIONS")
		http.Error(w, "Error: method not allowed", http.StatusMethodNotAllowed)
	}
}
