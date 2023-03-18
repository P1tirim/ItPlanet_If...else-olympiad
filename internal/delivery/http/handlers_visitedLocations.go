package http

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) AddVisitedLocation(c *fiber.Ctx) error {
	animalId, err := c.ParamsInt("animalId")
	if err != nil || animalId <= 0 {
		return c.SendStatus(400)
	}

	locationId, err := c.ParamsInt("pointId")
	if err != nil || locationId <= 0 {
		return c.SendStatus(400)
	}

	animal, err := h.services.GetAnimalById(animalId)
	if err != nil {
		log.Println(err)
		return c.SendStatus(500)
	}

	if animal.Id == 0 {
		return c.SendStatus(404)
	}

	if animal.LifeStatus == "DEAD" {
		return c.SendStatus(400)
	}

	locationId, _, _, err = h.services.GetLocations(locationId)
	if err != nil {
		log.Println(err)
		return c.SendStatus(500)
	}

	if locationId == 0 {
		return c.SendStatus(404)
	}

	if len(animal.VisitedLocations) == 0 && animal.ChippingLocationId == locationId {
		return c.SendStatus(400)
	}

	visitedLocations, err := h.services.GetVisitedLocationsAll(animalId)
	if err != nil {
		return c.SendStatus(500)
	}

	if len(animal.VisitedLocations) > 0 && locationId == visitedLocations[len(visitedLocations)-1].LocationId {
		return c.SendStatus(400)
	}

	visitedLocation, date, err := h.services.AddVisitedLocation(animalId, locationId)
	if err != nil {
		log.Println(err)
		return c.SendStatus(500)
	}

	c.Status(201)
	return c.JSON(fiber.Map{
		"id":                           visitedLocation,
		"dateTimeOfVisitLocationPoint": date.Format(time.RFC3339),
		"locationPointId":              locationId,
	})
}

func (h *Handlers) UpdateVisitedLocation(c *fiber.Ctx) error {
	animalId, err := c.ParamsInt("animalId")
	if err != nil || animalId <= 0 {
		return c.SendStatus(400)
	}

	var body map[string]interface{}

	err = json.Unmarshal(c.Body(), &body)
	if err != nil {
		return c.SendStatus(400)
	}

	visitedLocationId, ok := body["visitedLocationPointId"].(float64)
	if !ok || visitedLocationId <= 0 {
		return c.SendStatus(400)
	}

	locationId, ok := body["locationPointId"].(float64)
	if !ok || locationId <= 0 {
		return c.SendStatus(400)
	}

	locationIdCheck, _, _, err := h.services.GetLocations(int(locationId))
	if err != nil {
		return c.SendStatus(500)
	}

	if locationIdCheck == 0 {
		return c.SendStatus(404)
	}

	animal, err := h.services.GetAnimalById(animalId)
	if err != nil {
		return c.SendStatus(500)
	}

	if animal.Id == 0 {
		return c.SendStatus(404)
	}

	visitedLocations, err := h.services.GetVisitedLocationsAll(animalId)
	if err != nil {
		return c.SendStatus(500)
	}

	if len(visitedLocations) == 0 {
		return c.SendStatus(404)
	}

	var date time.Time

	for i := 0; i < len(visitedLocations); i++ {
		if visitedLocations[i].Id == int(visitedLocationId) {

			if i == 0 && animal.ChippingLocationId == int(locationId) {
				return c.SendStatus(400)
			}

			if visitedLocations[i].LocationId == int(locationId) {
				return c.SendStatus(400)
			}

			if i != 0 && visitedLocations[i-1].LocationId == int(locationId) {
				return c.SendStatus(400)
			}

			if i+1 < len(visitedLocations) && visitedLocations[i+1].LocationId == int(locationId) {
				return c.SendStatus(400)
			}

			date = visitedLocations[i].Date
		}
	}

	found, err := h.services.UpdateVisitedLocation(int(visitedLocationId), animalId, int(locationId))
	if err != nil {
		return c.SendStatus(500)
	}

	if !found {
		return c.SendStatus(404)
	}

	return c.JSON(fiber.Map{
		"id":                           visitedLocationId,
		"dateTimeOfVisitLocationPoint": date.Format(time.RFC3339),
		"locationPointId":              locationId,
	})
}

func (h *Handlers) DeleteVisitedLocation(c *fiber.Ctx) error {
	animalId, err := c.ParamsInt("animalId")
	if err != nil || animalId <= 0 {
		return c.SendStatus(400)
	}

	visitedLocation, err := c.ParamsInt("visitedPointId")
	if err != nil || visitedLocation <= 0 {
		return c.SendStatus(400)
	}

	animal, err := h.services.GetAnimalById(animalId)
	if err != nil {
		return c.SendStatus(500)
	}

	if animal.Id == 0 {
		return c.SendStatus(404)
	}

	found, err := h.services.DeleteVisitedLocation(visitedLocation, animal)
	if err != nil {
		return c.SendStatus(500)
	}

	if !found {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
