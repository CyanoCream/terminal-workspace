package handler

import (
	"github.com/gofiber/fiber/v2"
	"terminal/internal/user/domain"
	"terminal/internal/user/port"
	"terminal/pkg/jwt"
)

type userHandler struct {
	service port.UserService
}

func NewUserHandler(service port.UserService) port.UserHandler {
	return &userHandler{service: service}
}

func (h *userHandler) GetProfile(c *fiber.Ctx) error {
	claims := c.Locals("claims").(*jwt.Claims)
	user, err := h.service.GetProfile(claims.UserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get profile",
		})
	}
	return c.JSON(user)
}

func (h *userHandler) CreateCard(c *fiber.Ctx) error {
	claims := c.Locals("claims").(*jwt.Claims)

	var card domain.Card
	if err := c.BodyParser(&card); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	if err := h.service.CreateCard(claims.UserID, &card); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create card",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(card)
}

func (h *userHandler) TopUpCard(c *fiber.Ctx) error {
	claims := c.Locals("claims").(*jwt.Claims)

	var req struct {
		CardNumber string  `json:"card_number"`
		Amount     float64 `json:"amount"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	if err := h.service.TopUpCard(claims.UserID, req.CardNumber, req.Amount); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to top up card",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Card topped up successfully",
	})
}
