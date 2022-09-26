package handlers

import (
	episodedto "dumbflix/dto/episode"
	dto "dumbflix/dto/result"
	"dumbflix/models"
	"dumbflix/repository"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerEpisode struct {
	EpisodeRepository repository.EpisodeRepository
}

func HandlerEpisode(EpisodeRepository repository.EpisodeRepository) *handlerEpisode {
	return &handlerEpisode{EpisodeRepository}
}
func (h *handlerEpisode) FindAllEpisode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	episode, err := h.EpisodeRepository.FindAllEpisode()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: episode}
	json.NewEncoder(w).Encode(response)
}
func (h *handlerEpisode) GetEpisode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
  
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
  
	data, err := h.EpisodeRepository.GetEpisode(id)
	if err != nil {
	  w.WriteHeader(http.StatusBadRequest)
	  response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	  json.NewEncoder(w).Encode(response)
	  return
	}
  
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data:convertResponseEpisode(data)}
	json.NewEncoder(w).Encode(response)
  }
func (h *handlerEpisode) CreateEpisode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	filmId,_:=strconv.Atoi(r.FormValue("filmId"))

	request := episodedto.EpisodeReq{
		Title: r.FormValue("title"),
		ThumbnailFilm: r.FormValue("thumbnail"),
		LinkFilm: r.FormValue("LinkFilm"),
		Description: r.FormValue("description"),
		FilmID: filmId,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	episode := models.Episode{
		Title: request.Title,
		ThumbnailFilm : request.ThumbnailFilm,
		LinkFilm : request.LinkFilm,
		Description: request.Description,
		FilmID: request.FilmID,				       
	}

	episode, err = h.EpisodeRepository.CreateEpisode(episode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}
	episode,_ =h.EpisodeRepository.GetEpisode(episode.ID)
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: (episode)}
	json.NewEncoder(w).Encode(response)
}
func (h *handlerEpisode) UpdateEpisode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(episodedto.EpisodeReq)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	episode, err := h.EpisodeRepository.GetEpisode(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}


	data, err := h.EpisodeRepository.UpdateEpisode(episode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseEpisode(data)}
	json.NewEncoder(w).Encode(response)
}
func (h *handlerEpisode) DeleteEpisode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	episode, err := h.EpisodeRepository.GetEpisode(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.EpisodeRepository.DeleteEpisode(episode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseEpisode(data)}
	json.NewEncoder(w).Encode(response)
}




func convertResponseEpisode(u models.Episode) episodedto.EpisodeRes {
	return episodedto.EpisodeRes{
		Title: u.Title,
		ThumbnailFilm: u.ThumbnailFilm,
		LinkFilm: u.LinkFilm,
		Description : u.Description,
	}
}