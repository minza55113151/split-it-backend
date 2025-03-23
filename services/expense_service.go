package services

import (
	"errors"
	"split-it/models"
	"split-it/repositories"
)

type ExpenseService struct {
	expenseRepo *repositories.ExpenseRepository
}

func NewExpenseService(expenseRepo *repositories.ExpenseRepository) *ExpenseService {
	return &ExpenseService{
		expenseRepo: expenseRepo,
	}
}

func (s *ExpenseService) GetUserExpensesWithStatus(subID string, status string) ([]models.Expense, error) {
	return s.expenseRepo.GetUserExpensesWithStatus(subID, status)
}

func (s *ExpenseService) Create(expense *models.Expense) error {
	if err := s.validateExpense(expense); err != nil {
		return err
	}

	return s.expenseRepo.Create(expense)
}

func (s *ExpenseService) Update(id string, expense *models.Expense) error {
	if err := s.validateExpense(expense); err != nil {
		return err
	}

	oldExpense, err := s.expenseRepo.Get(id)
	if err != nil {
		return err
	}

	if expense.PayerSubID != oldExpense.PayerSubID && expense.PayerSubID != oldExpense.DebtorSubID ||
		expense.DebtorSubID != oldExpense.PayerSubID && expense.DebtorSubID != oldExpense.DebtorSubID {
		return errors.New("you can't change the payer or debtor of an expense")
	}

	return s.expenseRepo.Update(id, oldExpense)
}

func (s *ExpenseService) Delete(id string, subID string) error {
	return s.expenseRepo.Delete(id, subID)
}

func (s *ExpenseService) validateExpense(expense *models.Expense) error {
	if expense.PayerSubID == expense.DebtorSubID {
		return errors.New("you can't create an expense for yourself")
	}

	return nil
}
