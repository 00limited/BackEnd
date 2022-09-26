package models

type Episode struct {
	ID            int    `json:"id" gorm:"primary_key:auto_increment"`
	Title         string `json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm string `json:"thumbnail" form:"thumbnail"`
	LinkFilm      string `json:"LinkFilm" form:"LinkFilm" gorm:"type: varchar(255)"`
	Description   string `json:"description" gorm:"type:text" form:"description"`
	FilmID        int    `json:"filmid" form:"film_id"`
	Film          Film   `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}