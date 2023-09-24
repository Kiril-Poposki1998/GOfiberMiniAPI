package controller

import (
	"fmt"

	"github.com/Kiril-Poposki1998/GOfiberMiniAPI/model"

	"github.com/gofiber/fiber/v2"
)

var persons []*model.Person = nil

func SetPerson(c *fiber.Ctx) error {
	temp_person := new(model.Person)
	if err := c.BodyParser(temp_person); err != nil {
		fmt.Print(err.Error())
		return err
	}
	persons = append(persons, temp_person)
	return c.Status(fiber.StatusOK).SendString("Person created")
}

func GetPerson(c *fiber.Ctx) error {
	if persons != nil {
		return c.Render("index", fiber.Map{
			"Title":   "Persons",
			"Persons": persons,
		})
	}
	return c.Status(fiber.StatusOK).SendString("No person created")
}
