package main

import (
	"comingback/controller"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	//use logger middleware to monitor.
	app.Use(logger.New())

	//isolate the vendor and the user
	// app.Route("/", func(router fiber.Router) {
	// 	router.Get("/", func(ctx *fiber.Ctx) error {
	// 		ctx.Status(fiber.StatusNoContent).SendString("Nothing here!")

	// 		return fiber.NewError(fiber.StatusNoContent, "i said that nothing here!")
	// 	}).Name("error")
	// }, "nothing")
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(ctx.App().GetRoutes(false))
	}).Name("how to use the API")

	app.Route("/vendor", controller.VendorRoute, "vendor")

	log.Fatal(app.Listen(":3000"))
}

/*


use gin, fiber or echo as main framework, fast development
*/
