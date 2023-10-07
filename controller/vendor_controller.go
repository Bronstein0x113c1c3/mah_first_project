package controller

import (
	"comingback/model"

	"github.com/gofiber/fiber/v2"
)

func getAll(ctx *fiber.Ctx) error {
	return ctx.JSON(model.AllVendor())
}
func get(ctx *fiber.Ctx) error {
	
	return nil
}

func add(ctx *fiber.Ctx) error {
	
	return nil
}
func edit(ctx *fiber.Ctx) error {
	return nil
}
