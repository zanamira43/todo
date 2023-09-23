package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zanamira43/todo/database"
	"github.com/zanamira43/todo/models"
)

// get all Permission
func AllPermission(c *fiber.Ctx) error {
	var permissions []models.Permission
	database.DB.Find(&permissions)

	return c.JSON(permissions)
}

// create new permission
func CreatePermission(c *fiber.Ctx) error {
	var permission models.Permission

	if err := c.BodyParser(&permission); err != nil {
		return err
	}

	database.DB.Create(&permission)
	return c.JSON(permission)
}
