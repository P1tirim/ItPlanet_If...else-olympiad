package http

import (
	"api/pkg/models"
	"encoding/json"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) GetAnimalById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("animalId")
	if err != nil || id <= 0 {
		return c.SendStatus(400)
	}

	animal, err := h.services.GetAnimalById(id)
	if err != nil {
		return c.SendStatus(500)
	}

	if animal.Id == 0 {
		return c.SendStatus(404)
	}

	return c.JSON(animal)
}

func (h *Handlers) GetAnimalSearch(c *fiber.Ctx) error {
	var search models.AnimalSearch
	var err error

	search.StartDateTime = c.Query("startDateTime")
	if search.StartDateTime != "" {
		_, err := time.Parse(time.RFC3339, search.StartDateTime)
		if err != nil {
			return c.SendStatus(400)
		}
	}

	search.EndDateTime = c.Query("endDateTime")
	if search.EndDateTime != "" {
		_, err := time.Parse(time.RFC3339, search.EndDateTime)
		if err != nil {
			return c.SendStatus(400)
		}
	}

	search.ChipperId, err = strconv.Atoi(c.Query("chipperId", "0"))
	if err != nil || search.ChipperId < 0 {
		return c.SendStatus(400)
	}

	search.ChippingLocationId, err = strconv.Atoi(c.Query("chippingLocationId", "0"))
	if err != nil || search.ChippingLocationId < 0 {
		return c.SendStatus(400)
	}

	search.LifeStatus = c.Query("lifeStatus")
	if search.LifeStatus != "ALIVE" && search.LifeStatus != "DEAD" && search.LifeStatus != "" {
		return c.SendStatus(400)
	}

	search.Gender = c.Query("gender")
	if search.Gender != "MALE" && search.Gender != "FEMALE" && search.Gender != "OTHER" && search.Gender != "" {
		return c.SendStatus(400)
	}

	search.From, err = strconv.Atoi(c.Query("from", "0"))
	if err != nil || search.From < 0 {
		return c.SendStatus(400)
	}

	search.Size, err = strconv.Atoi(c.Query("size", "10"))
	if err != nil || search.Size <= 0 {
		return c.SendStatus(400)
	}

	animals, err := h.services.GetAnimalSearch(search)
	if err != nil {
		return c.SendStatus(500)
	}

	if len(animals) == 0 {
		animals = make([]models.Animal, 0)
	}

	return c.JSON(animals)
}

func (h *Handlers) AddAnimal(c *fiber.Ctx) error {
	var animal models.Animal

	err := json.Unmarshal(c.Body(), &animal)
	if err != nil {
		return c.SendStatus(400)
	}

	if animal.AnimalTypes == nil || len(animal.AnimalTypes) <= 0 {
		return c.SendStatus(400)
	}

	if !isArrayDistinct(animal.AnimalTypes) {
		return c.SendStatus(409)
	}

	for _, val := range animal.AnimalTypes {
		if val <= 0 {
			return c.SendStatus(400)
		}

		typeId, _, err := h.services.GetTypeOfAnimal(val)
		if err != nil {
			return c.SendStatus(500)
		}

		if typeId == 0 {
			return c.SendStatus(404)
		}
	}

	if animal.Weight <= 0 || animal.Length <= 0 || animal.Height <= 0 || animal.ChipperId <= 0 || animal.ChippingLocationId <= 0 {
		return c.SendStatus(400)
	}

	if animal.Gender != "MALE" && animal.Gender != "FEMALE" && animal.Gender != "OTHER" {
		return c.SendStatus(400)
	}

	account, err := h.services.GetAccountById(animal.ChipperId)
	if err != nil {
		return c.SendStatus(500)
	}

	if account.ID == 0 {
		return c.SendStatus(404)
	}

	locationId, _, _, err := h.services.GetLocations(animal.ChippingLocationId)
	if err != nil {
		return c.SendStatus(500)
	}

	if locationId == 0 {
		return c.SendStatus(404)
	}

	animal, err = h.services.AddAnimal(animal)
	if err != nil {
		return c.SendStatus(500)
	}

	c.Status(201)
	return c.JSON(animal)
}

func isArrayDistinct(arr []int) bool {
	var buf []int

	for _, val := range arr {
		for _, bufVal := range buf {
			if val == bufVal {
				return false
			}
			buf = append(buf, val)
		}
	}

	return true
}

func (h *Handlers) UpdateAnimal(c *fiber.Ctx) error {
	animalId, err := c.ParamsInt("animalId")
	if err != nil || animalId <= 0 {
		return c.SendStatus(400)
	}

	animalNew := models.Animal{}

	err = json.Unmarshal(c.Body(), &animalNew)
	if err != nil {
		return c.SendStatus(400)
	}

	if animalNew.Weight <= 0 || animalNew.Length <= 0 || animalNew.Height <= 0 || animalNew.ChipperId <= 0 || animalNew.ChippingLocationId <= 0 {
		return c.SendStatus(400)
	}

	if animalNew.Gender != "MALE" && animalNew.Gender != "FEMALE" && animalNew.Gender != "OTHER" {
		return c.SendStatus(400)
	}

	account, err := h.services.GetAccountById(animalNew.ChipperId)
	if err != nil {
		return c.SendStatus(500)
	}

	if account.ID == 0 {
		return c.SendStatus(404)
	}

	locationId, _, _, err := h.services.GetLocations(animalNew.ChippingLocationId)
	if err != nil {
		return c.SendStatus(500)
	}

	if locationId == 0 {
		return c.SendStatus(404)
	}

	animalOld, err := h.services.GetAnimalById(animalId)
	if err != nil {
		return c.SendStatus(500)
	}

	if animalOld.Id == 0 {
		return c.SendStatus(404)
	}

	if animalOld.LifeStatus == "DEAD" && animalNew.LifeStatus == "ALIVE" {
		return c.SendStatus(400)
	}

	visitedLocations, err := h.services.GetVisitedLocationsAll(animalOld.Id)
	if err != nil {
		return c.SendStatus(500)
	}

	if len(visitedLocations) > 0 && visitedLocations[0].LocationId == animalNew.ChippingLocationId {
		return c.SendStatus(400)
	}

	animalNew.Id = animalId

	err = h.services.UpdateAnimal(animalNew)
	if err != nil {
		return c.SendStatus(500)
	}

	if animalOld.LifeStatus == "ALIVE" && animalNew.LifeStatus == "DEAD" {
		err = h.services.UpdateDeathTimeAnimal(animalId)
		if err != nil {
			return c.SendStatus(500)
		}
	}

	animalNew, err = h.services.GetAnimalById(animalId)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(animalNew)
}

func (h *Handlers) DeleteAnimal(c *fiber.Ctx) error {
	animalId, err := c.ParamsInt("animalId")
	if err != nil || animalId <= 0 {
		return c.SendStatus(400)
	}

	visitedLocations, err := h.services.GetVisitedLocationsAll(animalId)
	if err != nil {
		return c.SendStatus(500)
	}

	if len(visitedLocations) > 0 {
		return c.SendStatus(400)
	}

	found, err := h.services.DeleteAnimal(animalId)
	if err != nil {
		return c.SendStatus(500)
	}

	if !found {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func (h *Handlers) AddTypeToAnimal(c *fiber.Ctx) error {
	animalId, err := c.ParamsInt("animalId")
	if err != nil || animalId <= 0 {
		return c.SendStatus(400)
	}

	typeId, err := c.ParamsInt("typeId")
	if err != nil || typeId <= 0 {
		return c.SendStatus(400)
	}

	animal, err := h.services.GetAnimalById(animalId)
	if err != nil {
		return c.SendStatus(500)
	}

	if animal.Id == 0 {
		return c.SendStatus(404)
	}

	typeId, _, err = h.services.GetTypeOfAnimal(typeId)
	if err != nil {
		return c.SendStatus(500)
	}

	if typeId == 0 {
		return c.SendStatus(404)
	}

	for _, animalType := range animal.AnimalTypes {
		if animalType == typeId {
			return c.SendStatus(409)
		}
	}

	err = h.services.AddTypeToAnimals(animalId, typeId)
	if err != nil {
		return c.SendStatus(500)
	}

	animal, err = h.services.GetAnimalById(animalId)
	if err != nil {
		return c.SendStatus(500)
	}

	c.Status(201)
	return c.JSON(animal)
}

func (h *Handlers) UpdateTypeOfAnimal(c *fiber.Ctx) error {
	animalId, err := c.ParamsInt("animalId")
	if err != nil || animalId <= 0 {
		return c.SendStatus(400)
	}

	var body map[string]interface{}

	err = json.Unmarshal(c.Body(), &body)
	if err != nil {
		return c.SendStatus(400)
	}

	oldType, ok := body["oldTypeId"].(float64)
	if !ok || oldType <= 0 {
		return c.SendStatus(400)
	}

	newType, ok := body["newTypeId"].(float64)
	if !ok || newType <= 0 {
		return c.SendStatus(400)
	}

	animal, err := h.services.GetAnimalById(animalId)
	if err != nil {
		return c.SendStatus(500)
	}

	if animal.Id == 0 {
		return c.SendStatus(404)
	}

	oldTypeCheck, _, err := h.services.GetTypeOfAnimal(int(oldType))
	if err != nil {
		return c.SendStatus(500)
	}
	if oldTypeCheck == 0 {
		return c.SendStatus(404)
	}

	newTypeCheck, _, err := h.services.GetTypeOfAnimal(int(newType))
	if err != nil {
		return c.SendStatus(500)
	}
	if newTypeCheck == 0 {
		return c.SendStatus(404)
	}

	haveOldType := false

	for _, animaltype := range animal.AnimalTypes {
		if animaltype == newTypeCheck {
			return c.SendStatus(409)
		}

		if animaltype == oldTypeCheck {
			haveOldType = true
		}
	}

	if !haveOldType {
		return c.SendStatus(404)
	}

	err = h.services.UpdateTypeOfAnimal(animalId, int(oldType), int(newType))
	if err != nil {
		c.SendStatus(500)
	}

	animal, err = h.services.GetAnimalById(animalId)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(animal)
}

func (h *Handlers) DeleteTypeOfAnimal(c *fiber.Ctx) error {
	animalId, err := c.ParamsInt("animalId")
	if err != nil || animalId <= 0 {
		return c.SendStatus(400)
	}

	typeId, err := c.ParamsInt("typeId")
	if err != nil || typeId <= 0 {
		return c.SendStatus(400)
	}

	animal, err := h.services.GetAnimalById(animalId)
	if err != nil {
		return c.SendStatus(500)
	}

	if animal.Id == 0 {
		return c.SendStatus(404)
	}

	types, err := h.services.GetTypesOfAnimal(animalId)
	if err != nil {
		return c.SendStatus(500)
	}

	if len(types) == 1 && types[0] == typeId {
		return c.SendStatus(400)
	}

	found, err := h.services.DeleteTypeOfAnimal(animalId, typeId)
	if err != nil {
		return c.SendStatus(500)
	}

	if !found {
		return c.SendStatus(404)
	}

	animal, err = h.services.GetAnimalById(animalId)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(animal)
}
