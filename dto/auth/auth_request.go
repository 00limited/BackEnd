package authdto

type AuthReq struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type RegisterReq struct {
	Fullname string `gorm:"type: varchar(255)" json:"fullname"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Gender   string `gorm:"type: varchar(225)" json:"gender" `
	Phone    string `gorm:"type: int" json:"phone"`
	Address  string `json:"address" gorm:"type: text"`
	Role     string `json:"role" gorm:"type: varchar(225)" `
}