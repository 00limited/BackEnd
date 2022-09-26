package episodedto

import "dumbflix/models"

type EpisodeRes struct {
	ID            int    `json:"id"`
	Title         string `json:"title" form:"title" validate:"required"`
	ThumbnailFilm string `json:"thumbnail" form:"thumbnail"`
	LinkFilm      string `json:"LinkFilm" form:"LinkFilm" gorm:"type: varchar(255)"`
	FilmID          int  `json:"filmid"`
	Film        models.Film `json:"film"`
	Description  string  `json:"description"`
}