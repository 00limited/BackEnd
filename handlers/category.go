package handlers

import (
	categorydto "dumbflix/dto/category"
	dto "dumbflix/dto/result"
	"dumbflix/models"
	"dumbflix/repository"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerCategory struct {
	CategoryRepository repository.CategoryRepository
}

func HandlerCategory(CategoryRepository repository.CategoryRepository) *handlerCategory {
	return &handlerCategory{CategoryRepository}
}
func (h *handlerCategory) FindAllCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	category, err := h.CategoryRepository.FindAllCategory()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	for i, cat := range category {
		for j, film := range cat.Film {
			category[i].Film[j].ThumbnailFilm = PathFile + film.ThumbnailFilm
		}
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: category}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCategory) GetCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	data, err := h.CategoryRepository.GetCategory(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	responseCategory := categorydto.CategoryRes{
		ID:   data.ID,
		Name: data.Name,
		Film: mapToCategory(data.Film),
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: responseCategory}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerCategory) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(categorydto.CategoryRes) //take pattern data submission
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	category, err := h.CategoryRepository.GetCategory(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Name != "" {
		category.Name = request.Name
	}

	data, err := h.CategoryRepository.UpdateCategory(category)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: (data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCategory) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	category, err := h.CategoryRepository.GetCategory(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	data, err := h.CategoryRepository.DeleteCategory(category)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: (data)}
	json.NewEncoder(w).Encode(response)
}
func (h *handlerCategory) CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(categorydto.CategoryRes)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// data form pattern submit to pattern entity db user
	category := models.Category{
		Name: request.Name,
	}

	data, err := h.CategoryRepository.CreateCategory(category)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: (data)}
	json.NewEncoder(w).Encode(response)
}

func mapToCategory(filmModels []models.Film) (mapCategoryResponse []categorydto.FilmResponseCategory) {
	for _, filmModel := range filmModels {
		filmResponse := categorydto.FilmResponseCategory{
			ID:            filmModel.ID,
			Title:         filmModel.Title,
			Description:   filmModel.Description,
			Year:          filmModel.Year,
			ThumbnailFilm: filmModel.ThumbnailFilm,
			LinkFilm:      filmModel.LinkFilm,
		}
		mapCategoryResponse = append(mapCategoryResponse, filmResponse)
	}
	return
}
