package models

import (
	"time"
)

type Film struct {
	ID            int        `json:"id" gorm:"primary_key:auto_increment"`
	Title         string     `json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm string     `json:"thumbnail" form:"thumbnail"`
	Year          string     `json:"year" form:"year" validate:"required"`
	Category      []Category `json:"category" gorm:"many2many:film_categories"`
	Description   string     `json:"description" gorm:"type:text" form:"description"`
	CreatedAt     time.Time  `json:"-"`
	UpdatedAt     time.Time  `json:"-"`
	Episode       []Episode  `json:"episode" `
	LinkFilm      string     `json:"linkfilm"`
}
