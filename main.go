package main

import (
	"log"
	"net/http"

	"github.com/faridEmilio/fiber-api/api/routes"
	internal "github.com/faridEmilio/fiber-api/internal/database"
	"github.com/gofiber/fiber/v2"
)

func InicializarApp(clienteHttp *http.Client) *fiber.App {

	internal.ConnectDb()
	app := fiber.New(fiber.Config{
		//Views: engine,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			var msg string
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
				msg = e.Message
			}

			if msg == "" {
				msg = "No se pudo procesar el llamado a la api: " + err.Error()
			}

			_ = ctx.Status(code).JSON(internalError{
				Message: msg,
			})

			return nil
		},
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("My first APi with Fiber")
	})

	// Prefijo comun a mis endpoints
	api := app.Group("/api/user")

	//Aca importo todos los endpoints de user
	routes.UserRoutes(api)

	return app

}

func main() {

	app := InicializarApp(http.DefaultClient)

	log.Fatal(app.Listen(":3000"))
}

type internalError struct {
	Message string `json:"message"`
}
