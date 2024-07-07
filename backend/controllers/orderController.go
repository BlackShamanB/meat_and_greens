package controllers

import (
	"main/database"
	"main/models"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func AllOrders(c fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(database.DB, &models.Order{}, page))
}
