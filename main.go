package main

import (
	"comingback/controller"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app:=fiber.New()
	app.Use(logger.New())
	app.Route("/vendor",controller.VendorRoute,"vendor")
	log.Fatal(app.Listen(:3000))
}

/*


use gin, fiber or echo as main framework, fast development
*/
