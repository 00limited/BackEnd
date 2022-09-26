package repository

import (
	"dumbflix/models"

	"gorm.io/gorm"
)

type AuthRepo interface {
	Login(email string) (models.User,error)
	Register(user models.User) (models.User,error)
	Getuser(ID int) (models.User, error)
}

func RepositoryAuth(db *gorm.DB) *repository{
	return &repository{db}
}

func (r *repository) Login(email string)(models.User, error) {

	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	// r.db.Raw("select * from Users where email = ?", email).Scan(&user)

	return user, err
}

func(r *repository) Register(user models.User) (models.User,error) {
	err := r.db.Create(&user).Error

	return user,err
}
func (r *repository) Getuser(ID int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, ID).Error

	return user, err
}