package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/middleware"
	"dumbflix/pkg/mysql"
	"dumbflix/repository"

	"github.com/gorilla/mux"
)

func Transaction(r *mux.Router) {
	transactionRepo := repository.RepositoryTransaction(mysql.DB)

	h := handlers.HandlerTransaction(transactionRepo)

	r.HandleFunc("/transaction",middleware.Auth(h.FindAllTransaction)).Methods("GET")
	r.HandleFunc("/transaction",middleware.Auth(h.CreateTransaction)).Methods("POST")
	// r.HandleFunc("/transaction/{id}",middleware.Auth(h.UpdateTransaction)).Methods("PATCH")
	r.HandleFunc("/transaction/{id}",h.DeleteTransaction).Methods("DELETE")
	r.HandleFunc("/notification", h.Notification).Methods("POST")

}