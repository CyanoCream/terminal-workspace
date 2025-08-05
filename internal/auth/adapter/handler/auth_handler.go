package handler

import (
	"github.com/gofiber/fiber/v2"
	"terminal/internal/auth/domain"
	"terminal/internal/auth/port"
)

type authHandler struct {
	service port.AuthService
}

func NewAuthHandler(service port.AuthService) port.AuthHandler {
	return &authHandler{service: service}
}

func (h *authHandler) Login(c *fiber.Ctx) error {
	var req domain.AuthRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	token, err := h.service.Login(req.Username, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Tambahkan message di response
	return c.JSON(fiber.Map{
		"message": "Login successful",
		"token":   token,
	})
}

func (h *authHandler) Register(c *fiber.Ctx) error {
	var user domain.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	user.Role = "user" // Default role
	if err := h.service.Register(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to register user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
	})
}
