package filmdto

type AddFilmRequest struct {
	Title         string `json:"title" form:"title" validate:"required"`
	ThumbnailFilm string `json:"thumbnail" form:"thumbnail"`
	Year          string `json:"year" form:"year" `
	CategoryID    []int  `json:"category_id" form:"category_id" gorm:"-"`
	Description   string `json:"description" gorm:"type:text" form:"description"`
	LinkFilm      string `json:"linkfilm" form:"linkfilm"`
}

type UpdateFilmRequest struct {
	Title         string `json:"title" form:"title" validate:"required"`
	ThumbnailFilm string `json:"thumbnail" form:"thumbnail"`
	Year          string `json:"year" form:"year" `
	CategoryID    []int  `json:"category_id" form:"category_id" gorm:"-"`
	Description   string `json:"description" gorm:"type:text" form:"description"`
	LinkFilm      string `json:"linkfilm" form:"linkfilm"`
}
