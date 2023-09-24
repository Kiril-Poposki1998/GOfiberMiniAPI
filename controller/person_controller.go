package controller

import (
	"fmt"
	"goapp/model"

	"github.com/gofiber/fiber/v2"
)

var person *model.Person = nil

func SetPerson(c *fiber.Ctx) error {
	person = new(model.Person)
	if err := c.BodyParser(person); err != nil {
		fmt.Print(err.Error())
		return err
	}
	return c.Status(fiber.StatusOK).SendString("Person created")
}

func GetPerson(c *fiber.Ctx) error {
	if person != nil {
		person_string := model.GetName(person) + " " + model.GetSurname(person)
		return c.Status(fiber.StatusOK).SendString(person_string)
	}
	return c.Status(fiber.StatusOK).SendString("No person created")
}
