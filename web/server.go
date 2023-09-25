package web

import (
	"anthnnygiang/api-template/app"
	"fmt"
	"net/http"
)

type Server struct {
	UserService    app.UserService
	SessionService app.SessionService
}

func (s *Server) HandleIndex(w http.ResponseWriter, r *http.Request) {
	//Forward slash matches everything
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Printf("%s %s\n", r.Method, r.URL.Path)

	// Common code for all requests...

	switch r.Method {
	case http.MethodGet:
		// Handle the GET request...
		fmt.Fprintf(w, "handling GET request\n")

	case http.MethodOptions:
		w.Header().Set("Allow", "GET, OPTIONS")
		w.WriteHeader(http.StatusNoContent)

	default:
		w.Header().Set("Allow", "GET, OPTIONS")
		http.Error(w, "error: method not allowed", http.StatusMethodNotAllowed)
	}
}
