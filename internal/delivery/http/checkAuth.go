package http

import (
	"api/pkg/models"
	"encoding/base64"
	"net/mail"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func (h *Handlers) CheckAuthForFun(c *fiber.Ctx) error {
	header := c.GetReqHeaders()
	auth := header["Authorization"]
	if auth == "" {
		return c.Next()
	}

	auth = strings.ReplaceAll(auth, "Basic ", "")

	data, err := base64.StdEncoding.DecodeString(auth)
	if err != nil {
		return c.SendStatus(401)
	}

	loginAndPass := strings.Split(string(data), ":")
	if len(loginAndPass) == 2 && h.services.CheckAuth(loginAndPass[0], loginAndPass[1]) {
		return c.Next()
	}

	return c.SendStatus(401)
}

func (h *Handlers) CheckAuth() func(*fiber.Ctx) error {
	config := basicauth.Config{Authorizer: h.services.CheckAuth}

	return basicauth.New(config)
}

func (h *Handlers) CheckOwnAccount(c *fiber.Ctx) error {
	id, err := c.ParamsInt("accountId")
	if err != nil || id <= 0 {
		return c.SendStatus(400)
	}

	email, ok := c.Locals("username").(string)
	if !ok {
		return c.SendStatus(401)
	}

	idCheck, err := h.services.GetIdAccountByEmail(email)
	if err != nil {
		return c.SendStatus(500)
	}

	if id != idCheck {
		return c.SendStatus(403)
	}

	return c.Next()
}

func isValidateAccountData(acc models.AccountReq) bool {
	if acc.FirstName == "" || !isValidString(acc.FirstName) {
		return false
	}

	if acc.LastName == "" || !isValidString(acc.LastName) {
		return false
	}

	if acc.Email == "" || !isValidString(acc.Email) {
		return false
	}

	_, err := mail.ParseAddress(acc.Email)
	if err != nil {
		return false
	}

	if acc.Password == "" || !isValidString(acc.Password) {
		return false
	}

	return true
}

func isValidString(s string) bool {
	for _, char := range s {
		if char == ' ' || char == '\t' || char == '\n' {
			return false
		}
	}
	return true
}
