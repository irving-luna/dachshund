package api

import (
	"fmt"
	"net/http"
	"positrace/models"

	"github.com/unrolled/render"
)

var (
	crender = render.New()
)

type APIServer struct {
	Port string
}

func NewAPIServer(port string) *APIServer {
	return &APIServer{
		Port: port,
	}
}

func (a *APIServer) Start() {

	// HTTP API server
	http.HandleFunc("/summary/stats", func(w http.ResponseWriter, r *http.Request) {
		// Retrieve and format client statistics
		// Return JSON response

		crender.JSON(w, http.StatusOK, models.Client{})
	})

	http.HandleFunc("/summary/active_sessions", func(w http.ResponseWriter, r *http.Request) {
		// Retrieve and format active sessions data
		// Return JSON response
		crender.JSON(w, http.StatusOK, []models.Client{	
		})
	})

	fmt.Printf("API serving in port: %s\n", a.Port)

	http.ListenAndServe(fmt.Sprintf(":%s", a.Port), nil)
}
