package controller

import "github.com/gofiber/fiber/v2"

//yeah, for isolating :)))
var VendorRoute = func(route fiber.Router) {
	route.Get("/get", getAll).Name(": get all the vendor")
	route.Get("/get/:id", get).Name(": get one vendor")
}
