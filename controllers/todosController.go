package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/zanamira43/todo/database"
	"github.com/zanamira43/todo/models"
)

// get all todos
func GetAllTodos(c *fiber.Ctx) error {
	var todos []models.Todo
	database.DB.Preload("TodoDate").Order("id desc").Find(&todos)
	return c.JSON(todos)
}

// Create Todo
func CreateTodo(c *fiber.Ctx) error {
	var todo models.Todo

	if err := c.BodyParser(&todo); err != nil {
		return err
	}

	database.DB.Create(&todo)
	return c.JSON(todo)
}

// get todo  by id
func GetSingleTodo(c *fiber.Ctx) error {
	// id, err := strconv.Atoi(c.Params("id"))

	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	var todo models.Todo
	database.DB.Preload("TodoDate").Where("id = ?", id).First(&todo)
	return c.JSON(todo)
}

// update single todo
func UpdateTodo(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	// id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	todo := models.Todo{
		Id: uint(id),
	}

	if err := c.BodyParser(&todo); err != nil {
		return err
	}

	database.DB.Model(&todo).Updates(todo)
	return c.JSON(todo)
}

// delete sintle todo
func DeleteTodo(c *fiber.Ctx) error {
	// id, err := strconv.Atoi(c.Params("id"))

	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	todo := models.Todo{
		Id: uint(id),
	}
	database.DB.Delete(&todo)
	return c.JSON(fiber.Map{
		"message": "Todo record successfully deleted",
	})
}
