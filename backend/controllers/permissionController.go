package controllers

import (
	"main/database"
	"main/models"

	"github.com/gofiber/fiber/v3"
)

func AllPermissions(c fiber.Ctx) error {
	var permissions []models.Permission

	database.DB.Find(&permissions)

	return c.JSON(permissions)
}
