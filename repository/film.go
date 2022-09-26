package repository

import (
	"dumbflix/models"

	"gorm.io/gorm"
)

type FilmRepository interface {
	FindAllFilm() ([]models.Film, error)
	GetFilm(ID int) (models.Film, error)
	CreateFilm(film models.Film) (models.Film, error)
	UpdateFilm(film models.Film) (models.Film, error)
	DeleteFilm(film models.Film, ID int) (models.Film, error)
	FindCategoriesById(CategoryID []int) ([]models.Category, error)
}

func RepositoryFilm(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAllFilm() ([]models.Film, error) {
	var movies []models.Film
	err := r.db.Preload("Category").Find(&movies).Error

	return movies, err
}
func (r *repository) GetFilm(ID int) (models.Film, error) {
	var film models.Film
	err := r.db.Preload("Category").First(&film, ID).Error

	return film, err
}
func (r *repository) CreateFilm(film models.Film) (models.Film, error) {
	err := r.db.Preload("Category").Create(&film).Error

	return film, err
}
func (r *repository) DeleteFilm(film models.Film, ID int) (models.Film, error) {
	err := r.db.Exec("SET FOREIGN_KEY_CHECKS=0").Delete(&film).Error

	return film, err
}
func (r *repository) UpdateFilm(film models.Film) (models.Film, error) {
	r.db.Model(&film).Association("Category").Replace(film.Category)

	err := r.db.Save(&film).Error

	return film, err
}

func (r *repository) FindCategoriesById(CategoryID []int) ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Find(&categories, CategoryID).Error

	return categories, err
}
