package handler

import (
	"github.com/gofiber/fiber/v2"
	"terminal/internal/terminal/domain"
	"terminal/internal/terminal/port"
)

type terminalHandler struct {
	service port.TerminalService
}

func NewTerminalHandler(service port.TerminalService) *terminalHandler {
	return &terminalHandler{service: service}
}

func (h *terminalHandler) GetAllTerminals(c *fiber.Ctx) error {
	terminals, err := h.service.GetAllTerminals()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get terminals",
		})
	}
	return c.JSON(terminals)
}

func (h *terminalHandler) CreateTerminal(c *fiber.Ctx) error {
	var terminal domain.Terminal
	if err := c.BodyParser(&terminal); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	if err := h.service.CreateTerminal(&terminal); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create terminal",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(terminal)
}

func (h *terminalHandler) AddGate(c *fiber.Ctx) error {
	var gate domain.Gate
	if err := c.BodyParser(&gate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	terminalID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid terminal ID",
		})
	}

	gate.TerminalID = uint(terminalID)
	if err := h.service.AddGate(&gate); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to add gate",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(gate)
}
func (h *terminalHandler) SetPricing(c *fiber.Ctx) error {
	var pricing domain.TerminalDistance
	if err := c.BodyParser(&pricing); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	if err := h.service.SetPricing(&pricing); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to set terminal pricing",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(pricing)
}
