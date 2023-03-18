package http

import (
	"api/pkg/models"
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) Registration(c *fiber.Ctx) error {

	header := c.GetReqHeaders()
	if header["Authorization"] != "" {
		return c.SendStatus(403)
	}

	var accountReq models.AccountReq

	err := json.Unmarshal(c.Body(), &accountReq)
	if err != nil {
		return c.SendStatus(400)
	}

	isValid := isValidateAccountData(accountReq)
	if !isValid {
		return c.SendStatus(400)
	}

	isDistinct, err := h.services.IsDistinctEmail(0, accountReq.Email)
	if err != nil {
		return c.SendStatus(500)
	}

	if !isDistinct {
		return c.SendStatus(409)
	}

	accountResp, err := h.services.Registration(accountReq)
	if err != nil {
		return c.SendStatus(500)
	}

	c = c.Status(201)
	return c.JSON(accountResp)
}

func (h *Handlers) GetAccountById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("accountId"))
	if err != nil || id <= 0 {
		return c.SendStatus(400)
	}

	acc, err := h.services.GetAccountById(id)
	if err != nil {
		return c.SendStatus(500)
	}

	if acc.ID == 0 {
		return c.SendStatus(404)
	}

	return c.JSON(acc)
}

func (h *Handlers) GetAccountSearch(c *fiber.Ctx) error {
	firstName := c.Query("firstName")

	lastName := c.Query("lastName")

	email := c.Query("email")

	from, err := strconv.Atoi(c.Query("from", "0"))
	if err != nil {
		return c.SendStatus(400)
	}

	size, err := strconv.Atoi(c.Query("size", "10"))
	if err != nil {
		return c.SendStatus(400)
	}

	if from < 0 || size <= 0 {
		return c.SendStatus(400)
	}

	accounts, err := h.services.GetAccountSearch(firstName, lastName, email, from, size)
	if err != nil {
		c.Status(500)
		return c.SendString(err.Error())
	}

	if len(accounts) == 0 {
		accounts = make([]models.AccountResp, 0)
	}

	return c.JSON(accounts)
}

func (h *Handlers) UpdateAccount(c *fiber.Ctx) error {
	var accountReq models.AccountReq

	err := json.Unmarshal(c.Body(), &accountReq)
	if err != nil {
		return c.SendStatus(400)
	}

	isValid := isValidateAccountData(accountReq)
	if !isValid {
		return c.SendStatus(400)
	}

	id, err := c.ParamsInt("accountId")
	if err != nil {
		return c.SendStatus(400)
	}

	isDistinct, err := h.services.IsDistinctEmail(id, accountReq.Email)
	if err != nil {
		return c.SendStatus(500)
	}

	if !isDistinct {
		return c.SendStatus(409)
	}

	err = h.services.UpdateAccountData(id, accountReq)
	if err != nil {
		return c.SendStatus(500)
	}

	acc, err := h.services.GetAccountById(id)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(acc)
}

func (h *Handlers) DeleteAccount(c *fiber.Ctx) error {
	id, err := c.ParamsInt("accountId")
	if err != nil || id <= 0 {
		return c.SendStatus(400)
	}

	b, err := h.services.IsAccountInAnimal(id)
	if err != nil {
		return c.SendStatus(500)
	}

	if b {
		return c.SendStatus(400)
	}

	b, err = h.services.DeleteAccount(id)
	if err != nil {
		return c.SendStatus(500)
	}

	if !b {
		return c.SendStatus(403)
	}

	return c.SendStatus(200)
}
