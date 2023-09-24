package controller

import (
	"goapp/model"

	"github.com/gofiber/fiber/v2"
)

func GetPerson(c *fiber.Ctx) error {
	temp := new(model.Person)
	model.SetName(temp, "kiril")
	model.SetSurname(temp, "Poposki")
	person_string := model.GetName(temp) + " " + model.GetSurname(temp)
	return c.Status(fiber.StatusOK).SendString(person_string)
}
