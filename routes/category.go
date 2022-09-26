package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/mysql"
	"dumbflix/repository"

	"github.com/gorilla/mux"
)

func Category(r *mux.Router) {
	CategoryRepo := repository.RepositoryCategory(mysql.DB)

	h := handlers.HandlerCategory(CategoryRepo)

	r.HandleFunc("/categories",h.FindAllCategory).Methods("GET")
	r.HandleFunc("/category",h.CreateCategory).Methods("POST")
	r.HandleFunc("/category/{id}",h.UpdateCategory).Methods("PATCH")
	r.HandleFunc("/category/{id}",h.DeleteCategory).Methods("DELETE")
	r.HandleFunc("/category/{id}",h.GetCategory).Methods("GET")

}