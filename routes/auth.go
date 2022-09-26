package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/middleware"
	"dumbflix/pkg/mysql"
	"dumbflix/repository"

	"github.com/gorilla/mux"
)

func Auth(r *mux.Router) {
	authRepo := repository.RepositoryAuth(mysql.DB)

	h := handlers.HandlerAuth(authRepo)

	r.HandleFunc("/login",h.Login).Methods("POST")
	r.HandleFunc("/register", h.Register).Methods("POST")
	r.HandleFunc("/check-auth", middleware.Auth(h.CheckAuth)).Methods("GET")
}