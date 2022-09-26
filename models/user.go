package models

type User struct {
	ID       int    `json:"id" gorm:"primary_key:auto_increment"`
	Name     string `json:"fullname" gorm:"type: varchar(255)"`
	Email    string `json:"email" gorm:"type:varchar(250)"`
	Password string `json:"password" gorm:"type: varchar(255)"`
	Gender   string `json:"gender" gorm:"type: varchar(255)"`
	Phone    string `json:"phone" gorm:"type: varchar(255)"`
	Address  string `json:"address" gorm:"type: text"`
	Role     string `json:"role" gorm:"type: varchar(255)"`
	Status   string `json:"status" gorm:"type:varchar(255)"`
}
