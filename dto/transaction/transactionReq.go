package transactiondto

import (
	"dumbflix/models"
	"time"
)

type AddTransaction struct {
	Attache   string          `json:"attache" form:"attache" gorm:"type: varchar(255)"`
	Status    string          `json:"status"  form:"status" gorm:"type:varchar(25)"`
}
type UpdateTransaction struct {
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	UserID int `json:"userid"`
	User  models.User `json:"user"`
	Attache   string          `json:"attache" form:"attache" gorm:"type: varchar(255)"`
	Status    string          `json:"status"  gorm:"type:varchar(25)"`
}

type DeleteTransaction struct{
	UserID int `json:"userid"`
	User  models.User `json:"user"`
}