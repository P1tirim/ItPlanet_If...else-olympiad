package http

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) GetTypeOfAnimal(c *fiber.Ctx) error {
	idType, err := c.ParamsInt("typeId")
	if err != nil || idType <= 0 {
		return c.SendStatus(400)
	}

	id, animalType, err := h.services.GetTypeOfAnimal(idType)
	if err != nil {
		return c.SendStatus(500)
	}

	if id == 0 {
		return c.SendStatus(404)
	}

	return c.JSON(fiber.Map{
		"id":   id,
		"type": animalType,
	})
}

func (h *Handlers) AddType(c *fiber.Ctx) error {
	var body map[string]string

	err := json.Unmarshal(c.Body(), &body)
	if err != nil {
		return c.SendStatus(400)
	}

	if !isValidString(body["type"]) || body["type"] == "" {
		return c.SendStatus(400)
	}

	distinct, err := h.services.IsDistinctType(body["type"])
	if err != nil {
		return c.SendStatus(500)
	}

	if !distinct {
		return c.SendStatus(409)
	}

	id, err := h.services.AddType(body["type"])
	if err != nil {
		return c.SendStatus(500)
	}

	c.SendStatus(201)
	return c.JSON(fiber.Map{
		"id":   id,
		"type": body["type"],
	})
}

func (h *Handlers) UpdateType(c *fiber.Ctx) error {
	id, err := c.ParamsInt("typeId")
	if err != nil || id <= 0 {
		return c.SendStatus(400)
	}

	var body map[string]string

	err = json.Unmarshal(c.Body(), &body)
	if err != nil {
		return c.SendStatus(400)
	}

	if !isValidString(body["type"]) || body["type"] == "" {
		return c.SendStatus(400)
	}

	distinct, err := h.services.IsDistinctType(body["type"])
	if err != nil {
		return c.SendStatus(500)
	}

	if !distinct {
		return c.SendStatus(409)
	}

	found, err := h.services.UpdateType(id, body["type"])
	if err != nil {
		return c.SendStatus(500)
	}

	if !found {
		return c.SendStatus(404)
	}

	return c.JSON(fiber.Map{
		"id":   id,
		"type": body["type"],
	})
}

func (h *Handlers) DeleteType(c *fiber.Ctx) error {
	id, err := c.ParamsInt("typeId")
	if err != nil || id <= 0 {
		return c.SendStatus(400)
	}

	inAnimals, err := h.services.IsTypeInAnimals(id)
	if err != nil {
		return c.SendStatus(500)
	}

	if inAnimals {
		return c.SendStatus(400)
	}

	found, err := h.services.DeleteType(id)
	if err != nil {
		return c.SendStatus(500)
	}

	if !found {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
