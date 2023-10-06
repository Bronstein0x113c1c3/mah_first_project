package controller

import "github.com/gofiber/fiber/v2"

var VendorRoute = func(route fiber.Router) {
	route.Get("/get").Name("")
	route.Post("/add")
}
