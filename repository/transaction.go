package repository

import (
	"dumbflix/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindAllTransaction() ([]models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	CreateTransaction(transactions models.Transaction) (models.Transaction, error)
	UpdateTransaction(status string, ID string) error
	DeleteTransaction(transaction models.Transaction) (models.Transaction, error)
	GetOneTransaction(ID string) (models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) FindAllTransaction() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("User").Find(&transactions).Error

	return transactions, err
}
func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transactions models.Transaction
	// not yet using category relation, cause this step doesnt Belong to Many
	err := r.db.Preload("User").First(&transactions, ID).Error

	return transactions, err
}
func (r *repository) CreateTransaction(transactions models.Transaction) (models.Transaction, error) {
	err := r.db.Exec("SET FOREIGN_KEY_CHECKS=0").Preload("User").Create(&transactions).Error

	return transactions, err
}
func (r *repository) DeleteTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Delete(&transaction).Error // Using Delete method

	return transaction, err
}
func (r *repository) UpdateTransaction(status string, ID string) error {
	var transaction models.Transaction
	r.db.Preload("User").First(&transaction, ID)

	// If is different & Status is "success" decrement product quantity
	if status != transaction.Status && status == "success" {
		var user models.User
		r.db.First(&user, transaction.User.ID)
		user.Status = "Active"
		r.db.Save(&user)
	}

	transaction.Status = status

	err := r.db.Save(&transaction).Error

	return err
}

// Create GetOneTransaction method here ...
func (r *repository) GetOneTransaction(ID string) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("User").First(&transaction, "id = ?", ID).Error

	return transaction, err
}
