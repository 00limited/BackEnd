package filmdto

import (
	"dumbflix/models"
)

type FilmResponse struct {
	ID            int               `json:"id" gorm:"primary_key:auto_increment"`
	Title         string            `json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm string            `json:"thumbnail" form:"thumbnail"`
	Year          string            `json:"year" form:"year"`
	Category      []models.Category `json:"category" gorm:"many2many:film_categories"`
	Description   string            `json:"description" gorm:"type:text" form:"description"`
	LinkFilm      string            `json:"linkfilm" form:"linkfilm"`
}
