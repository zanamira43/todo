package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/zanamira43/todo/database"
	"github.com/zanamira43/todo/models"
)

// get all Users
func AllUsers(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	return c.JSON(models.Paginate(database.DB, &models.User{}, page))
}

// create new user
func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	user.SetPassword("12345678")
	database.DB.Create(&user)
	return c.JSON(user)
}

// get single user
func GetUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	database.DB.Preload("Role").Find(&user)
	return c.JSON(user)
}

// update singel user
func UpdateUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	database.DB.Model(&user).Updates(user)
	return c.JSON(user)
}

// delete singel user
func DeleteUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}
	database.DB.Delete(&user)
	return nil
}
