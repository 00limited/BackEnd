package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/mysql"
	"dumbflix/repository"

	"github.com/gorilla/mux"
)

func Episode(r *mux.Router) {
	EpisodeRepository := repository.RepositoryEpisode(mysql.DB)

	h := handlers.HandlerEpisode(EpisodeRepository)

	r.HandleFunc("/film/{id_film}/episode",h.FindAllEpisode).Methods("GET")
	r.HandleFunc("/episode",h.CreateEpisode).Methods("POST")
	r.HandleFunc("/film/{id_film}/episode",h.GetEpisode).Methods("GET")
	r.HandleFunc("/episode/{id}",h.UpdateEpisode).Methods("PATCH")
	r.HandleFunc("/episode/{id}",h.DeleteEpisode).Methods("DELETE")
}