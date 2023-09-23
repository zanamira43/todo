package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/zanamira43/todo/database"
	"github.com/zanamira43/todo/models"
)

// get all todo date
func AllTodoDates(c *fiber.Ctx) error {
	var todoDates []models.TodoDate
	database.DB.Order("id desc").Find(&todoDates)
	return c.JSON(todoDates)
}

// Create Todo date
func CreateTodoDate(c *fiber.Ctx) error {
	var tododate models.TodoDate

	if err := c.BodyParser(&tododate); err != nil {
		return err
	}

	database.DB.Create(&tododate)
	return c.JSON(tododate)
}

// get todo date  by id
func GetSingleTodoDate(c *fiber.Ctx) error {
	// id, err := strconv.Atoi(c.Params("id"))
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	var tododate models.TodoDate
	database.DB.Where("id = ?", id).First(&tododate)
	return c.JSON(tododate)
}

// update single todo date
func UpdateTodoDate(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	// id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	tododate := models.TodoDate{
		ID: uint(id),
	}

	if err := c.BodyParser(&tododate); err != nil {
		return err
	}
	database.DB.Model(&tododate).Updates(tododate)
	return c.JSON(tododate)
}

// delete sintle todo date
func DeleteTodoDate(c *fiber.Ctx) error {
	// id, err := strconv.Atoi(c.Params("id"))

	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	tododate := models.TodoDate{
		ID: uint(id),
	}
	database.DB.Delete(&tododate)
	return nil
}

// get todo date by  using date
func GetSingleTodoDateByDate(c *fiber.Ctx) error {

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var todoDate models.TodoDate
	database.DB.Preload("Todo").Where("date = ? ", data["date"]).First(&todoDate)

	return c.JSON(todoDate)
}
