package transactiondto

import (
	"dumbflix/models"
	"time"
)

type Transaction struct {
	ID        int         `json:"id" gorm:"primary_key:auto_increment"`
	CreatedAt time.Time   `json:"-"`
	UpdatedAt int         `json:"-"`
	UserID    int         `json:"user_id"`
	User      models.User `json:"user"`
	Attache   string      `json:"attache" form:"attache" gorm:"type: varchar(255)"`
	Status    string      `json:"status"  gorm:"type:varchar(25)"`
}
