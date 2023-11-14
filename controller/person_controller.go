package controller

import (
	"fmt"

	"github.com/Kiril-Poposki1998/GOfiberMiniAPI/database"
	"github.com/Kiril-Poposki1998/GOfiberMiniAPI/middleware"
	"github.com/Kiril-Poposki1998/GOfiberMiniAPI/model"

	"github.com/gofiber/fiber/v2"
)

func SetPerson(c *fiber.Ctx) error {
	temp_person := new(model.Person)
	if err := c.BodyParser(temp_person); err != nil {
		fmt.Print(err.Error())
		return err
	}
	database.Database.Create(temp_person)
	return c.Status(fiber.StatusOK).SendString("Person created")
}

func GetPersons(c *fiber.Ctx) error {
	var persons []model.Person
	database.Database.Find(&persons)
	if len(persons) != 0 {
		return c.Render("index", fiber.Map{
			"Title":   "Persons",
			"Persons": persons,
		})
	}
	return c.Status(fiber.StatusOK).SendString("No person created")
}

func GetPerson(c *fiber.Ctx) error {
	if !middleware.CheckUser(c) {
		return c.Status(fiber.StatusForbidden).SendString("You are not authenticated to view users individually")
	}

	person := new(model.Person)
	id := c.Params("id")

	query := database.Database.Find(&person, id)

	if query.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).SendString("Person by id was not found")
	}

	person_json := struct {
		Id      uint
		Name    string
		Surname string
	}{person.Id, person.Name, person.Surname}

	return c.JSON(person_json)
}

func DeletePerson(c *fiber.Ctx) error {

	id := c.Params("id")

	query := database.Database.Delete(&model.Person{}, id)

	if query.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).SendString("Person by id was not found")
	}
	return c.Status(fiber.StatusOK).SendString("Person was deleted")
}

func UpdatePerson(c *fiber.Ctx) error {
	new_person := new(model.Person)
	id := c.Params("id")
	if err := c.BodyParser(new_person); err != nil {
		fmt.Print(err.Error())
		return err
	}
	database.Database.Where("id = ?", id).Updates(&new_person)
	return c.Status(fiber.StatusOK).SendString("Person update")
}
