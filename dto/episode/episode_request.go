package episodedto

type EpisodeReq struct {
	Title         string `json:"title" form:"title" validate:"required"`
	ThumbnailFilm string `json:"thumbnail" form:"thumbnail"`
	LinkFilm      string `json:"LinkFilm" form:"LinkFilm" gorm:"type: varchar(255)"`
	FilmID        int    `json:"filmid"`
	Description   string `json:"description"`
}