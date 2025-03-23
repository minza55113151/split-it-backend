package handlers

import (
	"split-it/models"
	"split-it/services"

	"github.com/gofiber/fiber/v2"
)

type ExpenseHandler struct {
	expenseService *services.ExpenseService
}

func NewExpenseHandler(expenseService *services.ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{expenseService: expenseService}
}

// @Summary Get user expenses with status
// @Description Get user expenses with status
// @Tags expenses
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param status path string true "Expense status"
// @Success 200 {array} models.Expense
// @Failure 500 {string} string
// @Router /expenses/{status} [get]
func (h *ExpenseHandler) HandleGetUserExpensesWithStatus(c *fiber.Ctx) error {
	subID := c.Locals(models.SubIDContextKey).(string)
	status := c.Params("status")

	expenses, err := h.expenseService.GetUserExpensesWithStatus(subID, status)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error()) // TODO: should return a more generic error message
	}

	return c.Status(fiber.StatusOK).JSON(expenses)
}

// @Summary Create expense
// @Description Create expense
// @Tags expenses
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param expense body models.Expense true "Expense"
// @Success 201 {string} string
// @Failure 500 {string} string
// @Router /expenses [post]
func (h *ExpenseHandler) HandleCreateExpense(c *fiber.Ctx) error {
	subID := c.Locals(models.SubIDContextKey).(string)

	expense := new(models.Expense)
	if err := c.BodyParser(expense); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if subID != expense.PayerSubID && subID != expense.DebtorSubID {
		return c.Status(fiber.StatusForbidden).SendString("You can only create an expense for yourself")
	}

	err := h.expenseService.Create(expense)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error()) // TODO: should return a more generic error message
	}

	return c.SendStatus(fiber.StatusCreated)
}

// @Summary Update expense
// @Description Update expense
// @Tags expenses
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Expense ID"
// @Param expense body models.Expense true "Expense"
// @Success 200 {string} string
// @Failure 500 {string} string
// @Router /expenses/{id} [put]
func (h *ExpenseHandler) HandleUpdateExpense(c *fiber.Ctx) error {
	subID := c.Locals(models.SubIDContextKey).(string)
	id := c.Params("id")

	expense := new(models.Expense)
	if err := c.BodyParser(expense); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if subID != expense.PayerSubID && subID != expense.DebtorSubID {
		return c.Status(fiber.StatusForbidden).SendString("You can only update an expense for yourself")
	}

	err := h.expenseService.Update(id, expense)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error()) // TODO: should return a more generic error message
	}

	return c.SendStatus(fiber.StatusOK)
}

// @Summary Delete expense
// @Description Delete expense
// @Tags expenses
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Expense ID"
// @Success 204 {string} string
// @Failure 500 {string} string
// @Router /expenses/{id} [delete]
func (h *ExpenseHandler) HandleDeleteExpense(c *fiber.Ctx) error {
	subID := c.Locals(models.SubIDContextKey).(string)
	id := c.Params("id")

	err := h.expenseService.Delete(id, subID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error()) // TODO: should return a more generic error message
	}

	return c.SendStatus(fiber.StatusOK)
}
