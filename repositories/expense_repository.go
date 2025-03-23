package repositories

import (
	"split-it/models"

	"gorm.io/gorm"
)

type ExpenseRepository struct {
	db *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) *ExpenseRepository {
	return &ExpenseRepository{db: db}
}

func (r *ExpenseRepository) Get(id string) (*models.Expense, error) {
	var expense models.Expense

	err := r.db.First(&expense, id).Error

	return &expense, err
}

func (r *ExpenseRepository) GetUserExpensesWithStatus(subID string, status string) ([]models.Expense, error) {
	var expenses []models.Expense

	err := r.db.Where("(payer_sub_id = ? OR debtor_sub_id = ?) AND status = ?", subID, subID, status).Find(&expenses).Error

	return expenses, err
}

func (r *ExpenseRepository) Create(expense *models.Expense) error {
	return r.db.Create(expense).Error
}

func (r *ExpenseRepository) Update(id string, expoense *models.Expense) error {
	return r.db.Model(&models.Expense{}).Where("id = ?", id).Updates(expoense).Error
}

func (r *ExpenseRepository) Delete(id string, subID string) error {
	return r.db.Where("id = ? AND (payer_sub_id = ? OR debtor_sub_id = ?)", id, subID, subID).Delete(&models.Expense{}).Error
}
