package http

import (
	"api/pkg/models"
	"encoding/json"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) GetLocations(c *fiber.Ctx) error {
	idPoint, err := c.ParamsInt("pointId")
	if err != nil || idPoint <= 0 {
		return c.SendStatus(400)
	}

	id, latitude, longitude, err := h.services.GetLocations(idPoint)
	if err != nil {
		return c.SendStatus(500)
	}

	if id == 0 {
		return c.SendStatus(404)
	}

	return c.JSON(fiber.Map{
		"id":        id,
		"latitude":  latitude,
		"longitude": longitude,
	})
}

func (h *Handlers) GetVisitedLocations(c *fiber.Ctx) error {
	idAnimal, err := c.ParamsInt("animalId")
	if err != nil || idAnimal <= 0 {
		return c.SendStatus(400)
	}

	startDate := c.Query("startDateTime")
	if startDate != "" {
		t, err := time.Parse(time.RFC3339, startDate)
		if err != nil {
			return c.SendStatus(400)
		}

		startDate = t.String()
	}

	endDate := c.Query("endDateTime")
	if endDate != "" {
		t, err := time.Parse(time.RFC3339, endDate)
		if err != nil {
			return c.SendStatus(400)
		}

		endDate = t.String()
	}

	from, err := strconv.Atoi(c.Query("from", "0"))
	if err != nil || from < 0 {
		return c.SendStatus(400)
	}

	size, err := strconv.Atoi(c.Query("size", "10"))
	if err != nil || size <= 0 {
		return c.SendStatus(400)
	}

	locations, err := h.services.GetVisitedLocations(idAnimal, startDate, endDate, from, size)
	if err != nil {
		return c.SendStatus(500)
	}

	if len(locations) == 0 {
		return c.SendStatus(404)
	}

	return c.JSON(locations)
}

func (h *Handlers) AddLocation(c *fiber.Ctx) error {
	var location models.Location

	err := json.Unmarshal(c.Body(), &location)
	if err != nil {
		return c.SendStatus(400)
	}

	if string(location.Latitude) == "null" || string(location.Longitude) == "null" {
		return c.SendStatus(400)
	}

	latitude, err := strconv.ParseFloat(string(location.Latitude), 64)
	if err != nil {
		return c.SendStatus(400)
	}

	longitude, err := strconv.ParseFloat(string(location.Longitude), 64)
	if err != nil {
		return c.SendStatus(400)
	}

	valid := isValidCoordinates(latitude, longitude)
	if !valid {
		return c.SendStatus(400)
	}

	distinct, err := h.services.IsDistintcCoordinates(float64(latitude), float64(longitude))
	if err != nil {
		return c.SendStatus(500)
	}

	if !distinct {
		return c.SendStatus(409)
	}

	id, err := h.services.AddLocation(float64(latitude), float64(longitude))
	if err != nil {
		return c.SendStatus(500)
	}

	c.Status(201)
	return c.JSON(fiber.Map{
		"id":        id,
		"latitude":  latitude,
		"longitude": longitude,
	})
}

func (h *Handlers) UpdateLocations(c *fiber.Ctx) error {
	id, err := c.ParamsInt("pointId")
	if err != nil || id <= 0 {
		return c.SendStatus(400)
	}

	var location models.Location

	err = json.Unmarshal(c.Body(), &location)
	if err != nil {
		return c.SendStatus(400)
	}

	if string(location.Latitude) == "null" || string(location.Longitude) == "null" {
		return c.SendStatus(400)
	}

	latitude, err := strconv.ParseFloat(string(location.Latitude), 64)
	if err != nil {
		return c.SendStatus(400)
	}

	longitude, err := strconv.ParseFloat(string(location.Longitude), 64)
	if err != nil {
		return c.SendStatus(400)
	}

	valid := isValidCoordinates(latitude, longitude)
	if !valid {
		return c.SendStatus(400)
	}

	distinct, err := h.services.IsDistintcCoordinates(latitude, longitude)
	if err != nil {
		return c.SendStatus(500)
	}

	if !distinct {
		return c.SendStatus(409)
	}

	found, err := h.services.UpdateLocations(id, latitude, longitude)
	if err != nil {
		return c.SendStatus(500)
	}

	if !found {
		return c.SendStatus(404)
	}

	return c.JSON(fiber.Map{
		"id":        id,
		"latitude":  latitude,
		"longitude": longitude,
	})
}

func isValidCoordinates(latitude, longitude float64) bool {

	if latitude < -90 || latitude > 90 {
		return false
	}

	if longitude < -180 || longitude > 180 {
		return false
	}

	return true
}

func (h *Handlers) DeleteLocation(c *fiber.Ctx) error {
	id, err := c.ParamsInt("pointId")
	if err != nil || id <= 0 {
		return c.SendStatus(400)
	}

	inAnimals, err := h.services.IsLocationInAnimals(id)
	if err != nil {
		return c.SendStatus(500)
	}

	if inAnimals {
		return c.SendStatus(400)
	}

	inVisitedLocation, err := h.services.IsLocationInVisitedLocation(id)
	if err != nil {
		return c.SendStatus(500)
	}

	if inVisitedLocation {
		return c.SendStatus(400)
	}

	found, err := h.services.DeleteLocaions(id)
	if err != nil {
		return c.SendStatus(500)
	}

	if !found {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
