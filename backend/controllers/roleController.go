package controllers

import (
	"main/database"
	"main/models"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func AllRoles(c fiber.Ctx) error {
	var roles []models.Role

	database.DB.Find(&roles)

	return c.JSON(roles)
}

func CreateRole(c fiber.Ctx) error {
	var roleDTO fiber.Map

	if err := c.Bind().Body(&roleDTO); err != nil {
		return err
	}

	list := roleDTO["permissions"].([]interface{})

	permissions := make([]models.Permission, len(list))

	for i, perpermissionId := range list {
		id, _ := strconv.Atoi(perpermissionId.(string))

		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}

	role := models.Role{
		Name:        roleDTO["name"].(string),
		Permissions: permissions,
	}

	database.DB.Create(&role)

	return c.JSON(role)
}

func GetRole(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	database.DB.Preload("Permissions").Find(&role)

	return c.JSON(role)
}

func UpdateRole(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var roleDTO fiber.Map

	if err := c.Bind().Body(&roleDTO); err != nil {
		return err
	}

	list := roleDTO["permissions"].([]interface{})

	permissions := make([]models.Permission, len(list))

	for i, perpermissionId := range list {
		id, _ := strconv.Atoi(perpermissionId.(string))

		permissions[i] = models.Permission{
			Id: uint(id),
		}
	}

	var result interface{}

	database.DB.Table("role_permissions").Where("role_id", id).Delete(result)

	role := models.Role{
		Id:          uint(id),
		Name:        roleDTO["name"].(string),
		Permissions: permissions,
	}
	database.DB.Model(&role).Updates(role)

	return c.JSON(role)
}

func DeleteRole(c fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}
	database.DB.Delete(&role)

	return nil
}
