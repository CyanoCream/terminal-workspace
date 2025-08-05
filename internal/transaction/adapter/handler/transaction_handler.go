package handler

import (
	"github.com/gofiber/fiber/v2"
	"terminal/internal/transaction/domain"
	"terminal/internal/transaction/port"
	"terminal/pkg/jwt"
)

type transactionHandler struct {
	service port.TransactionService
}

func NewTransactionHandler(service port.TransactionService) port.TransactionHandler {
	return &transactionHandler{service: service}
}

func (h *transactionHandler) CheckIn(c *fiber.Ctx) error {
	var req struct {
		CardNumber string `json:"card_number"`
		GateID     uint   `json:"gate_id"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	if err := h.service.CheckIn(req.CardNumber, req.GateID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Checkin successful",
	})
}

func (h *transactionHandler) CheckOut(c *fiber.Ctx) error {
	var req struct {
		CardNumber string `json:"card_number"`
		GateID     uint   `json:"gate_id"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	if err := h.service.CheckOut(req.CardNumber, req.GateID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Checkout successful",
	})
}

func (h *transactionHandler) SyncTransactions(c *fiber.Ctx) error {
	var transactions []domain.OfflineTransaction
	if err := c.BodyParser(&transactions); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	if err := h.service.SyncTransactions(transactions); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Transactions synced successfully",
	})
}

func (h *transactionHandler) GetTransactionHistory(c *fiber.Ctx) error {
	claims := c.Locals("claims").(*jwt.Claims)
	transactions, err := h.service.GetTransactionHistory(claims.UserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get transaction history",
		})
	}
	return c.JSON(transactions)
}
