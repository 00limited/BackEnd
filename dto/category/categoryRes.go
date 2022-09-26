package categorydto

type CategoryRes struct {
	ID   int                    `json:"id" gorm:"primary_key:auto_increment"`
	Name string                 `json:"name" form:"name" gorm:"type: varchar(255)"`
	Film []FilmResponseCategory `json:"film"`
}

type CategoryDel struct {
	ID int `json:"id" gorm:"primary_key:auto_increment"`
}

type FilmResponseCategory struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	ThumbnailFilm string `json:"thumbnail"`
	Year          string `json:"year"`
	Description   string `json:"description"`
	LinkFilm      string `json:"linkfilm"`
}
