package models

import "time"

type Transaction struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	CreatedAt time.Time `json:"timeStart"`
	UpdatedAt int       `json:"timeEnd"`
	UserID    int       `json:"userid" `
	User      User      `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Attache   string    `json:"attache" form:"atttache" gorm:"type: varchar(255)"`
	Status    string    `json:"status"  gorm:"type:varchar(25)"`
	Price     int       `json:"price"`
}
