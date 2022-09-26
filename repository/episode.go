package repository

import (
	"dumbflix/models"

	"gorm.io/gorm"
)

type EpisodeRepository interface {
	FindAllEpisode() ([]models.Episode, error)
	GetEpisode(ID int) (models.Episode, error)
	CreateEpisode(episode models.Episode) (models.Episode, error)
	UpdateEpisode(episode models.Episode) (models.Episode, error)
	DeleteEpisode(episode models.Episode) (models.Episode, error)
	GetFilmID(FilmID int) (models.Film, error)
}
func RepositoryEpisode(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) FindAllEpisode() ([]models.Episode, error) {
	var episode []models.Episode
	err := r.db.Preload("Film").Preload("Category").Find(&episode).Error

	return episode, err
}
func (r *repository) GetEpisode(ID int) (models.Episode, error) {
	var episode models.Episode
	// not yet using category relation, cause this step doesnt Belong to Many
	err := r.db.Preload("Film").Preload("Category").First(&episode, ID).Error
  
	return episode, err
  }
func (r *repository) CreateEpisode(episode models.Episode) (models.Episode, error) {
	err := r.db.Exec("SET FOREIGN_KEY_CHECKS=0").Create(&episode).Error
  
	return episode, err
  }
func (r *repository) UpdateEpisode(episode models.Episode) (models.Episode, error) {
	err := r.db.Preload("Film").Preload("Category").Save(&episode).Error
	
	return episode, err
}
func (r *repository) DeleteEpisode(episode models.Episode) (models.Episode, error) {
	err := r.db.Delete(&episode).Error // Using Delete method
  
	return episode, err
  }
  func (r *repository) GetFilmID(FilmID int) (models.Film, error) {
	var film models.Film
	err := r.db.Find(&film, FilmID).Error

	return film, err
}