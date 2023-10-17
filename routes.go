package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"gobase.com/base/pkg/api"
)

func initializeRoutes(r *mux.Router, s *services) {
	r.HandleFunc("/api/v0/users", api.RegisterUser(s.registering)).Methods("POST")
	r.HandleFunc("/api/test", test).Methods("POST")
}

func test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Go base project is running"))
}
