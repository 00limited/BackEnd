package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/middleware"
	"dumbflix/pkg/mysql"
	"dumbflix/repository"

	"github.com/gorilla/mux"
)

func Film(r *mux.Router) {
	filmRepo := repository.RepositoryFilm(mysql.DB)

	h := handlers.HandlerFilm(filmRepo)

	r.HandleFunc("/film", h.FindAllFilm).Methods("GET")
	r.HandleFunc("/film", middleware.Auth(middleware.UploadFile(h.CreateFilm))).Methods("POST")
	r.HandleFunc("/film/{id}", middleware.Auth(middleware.UploadFile(h.UpdateFilm))).Methods("PATCH")
	r.HandleFunc("/film/{id}", h.DeleteFilm).Methods("DELETE")
	r.HandleFunc("/film/{id}", middleware.Auth(h.GetFilm)).Methods("GET")

}
