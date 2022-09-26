package models

import "time"

type Category struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	Name      string    `json:"name"`
	Film      []Film    `json:"film" gorm:"many2many:film_categories"`
	FilmID    int       `json:"film_id"`
	CreatedAt time.Time `json:""`
	UpdatedAt time.Time `json:"-"`
}
