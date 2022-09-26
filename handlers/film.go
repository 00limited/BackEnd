package handlers

import (
	filmdto "dumbflix/dto/film"
	dto "dumbflix/dto/result"
	"dumbflix/models"
	"dumbflix/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"context"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerFilm struct {
	FilmRepository repository.FilmRepository
}

// var PathFile = "http://localhost:5000/upload/"

func HandlerFilm(FilmRepository repository.FilmRepository) *handlerFilm {
	return &handlerFilm{FilmRepository}
}
func (h *handlerFilm) FindAllFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// authorizationHeader := r.Header.Get("Authorization")
	// // if !strings.Contains(authorizationHeader, "Bearer") {
	// // 	http.Error(w, "Invalid token", http.StatusBadRequest)
	// // 	return
	// // }

	// tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
	// _, err := jwtToken.VerifyToken(tokenString)

	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }

	film, err := h.FilmRepository.FindAllFilm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	// // Create Embed Path File on Image & Video property here ...
	// for i, p := range film {
	// 	film[i].ThumbnailFilm = PathFile + p.ThumbnailFilm
	// }

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: film}
	json.NewEncoder(w).Encode(response)
}
func (h *handlerFilm) GetFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	film, err := h.FilmRepository.GetFilm(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// film.ThumbnailFilm = PathFile + film.ThumbnailFilm

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: (film)}
	json.NewEncoder(w).Encode(response)
}
func (h *handlerFilm) UpdateFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get product id
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	// //get data user token
	// userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	// userId := int(userInfo["id"].(float64))

	//get image filename
	dataContex := r.Context().Value("dataFile")
	filename := dataContex.(string)

	var categoriesId []int
	for _, r := range r.FormValue("categoryId") {
		if int(r-'0') >= 0 {
			categoriesId = append(categoriesId, int(r-'0'))
		}
	}

	request := filmdto.UpdateFilmRequest{
		Title:       r.FormValue("title"),
		Year:        r.FormValue("year"),
		Description: r.FormValue("description"),
		LinkFilm:    r.FormValue("linkFilm"),
		CategoryID:  categoriesId,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusContinue)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	//Get all category data by id []
	var category []models.Category
	if len(categoriesId) != 0 {
		category, _ = h.FilmRepository.FindCategoriesById(categoriesId)
	}
	film, _ := h.FilmRepository.GetFilm(id)

	film.Title = request.Title
	film.Description = request.Description
	film.Year = request.Year
	film.LinkFilm = request.LinkFilm
	film.Category = category

	if filename != "false" {
		film.ThumbnailFilm = filename
	}

	data, err := h.FilmRepository.UpdateFilm(film)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}
func (h *handlerFilm) DeleteFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	user, err := h.FilmRepository.GetFilm(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	data, err := h.FilmRepository.DeleteFilm(user, id)
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
func (h *handlerFilm) CreateFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Get dataFile from midleware and store to filename variable here ...
	fmt.Println("filename 1")

	dataContex := r.Context().Value("dataFile")
	filepath := dataContex.(string)

	var categoriesId []int
	for _, r := range r.FormValue("category_id") {
		if int(r-'0') >= 0 {
			categoriesId = append(categoriesId, int(r-'0'))
		}
	}
	// data form pattern submit to pattern entity db user
	request := filmdto.AddFilmRequest{
		Title:         r.FormValue("title"),
		ThumbnailFilm: r.FormValue("thumbnail"),
		Year:          r.FormValue("year"),
		CategoryID:    categoriesId,
		Description:   r.FormValue("description"),
		LinkFilm:      r.FormValue("linkfilm"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	category, _ := h.FilmRepository.FindCategoriesById(categoriesId)

	// Declare Context Background, Cloud Name, API Key, API Secret ...
	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	// Add your Cloudinary credentials ...
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "dumbways"})

	if err != nil {
		fmt.Println(err.Error())
	}

	film := models.Film{
		Title:         request.Title,
		ThumbnailFilm: resp.SecureURL,
		Year:          request.Year,
		Category:      category,
		Description:   request.Description,
		LinkFilm:      request.LinkFilm,
	}

	film, err = h.FilmRepository.CreateFilm(film)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	film, _ = h.FilmRepository.GetFilm(film.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: film}
	json.NewEncoder(w).Encode(response)
}
