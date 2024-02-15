package main

import (
	"log"
	"net/http"

	"github.com/faridEmilio/fiber-api/api/routes"
	"github.com/faridEmilio/fiber-api/internal/database"
	"github.com/faridEmilio/fiber-api/pkg/domains/user"
	"github.com/gofiber/fiber/v2"
)

func InicializarApp(clienteHttp *http.Client, clienteSql *database.DbInstance) *fiber.App {

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

	//Inicializo los repositorios con la base de datos
	//***** Repositories *****//
	userRepository := user.NewUserRepository(clienteSql)

	//Inicializo los servicios con el repositorio instanciado anteriormente
	//***** Services *****//
	userService := user.NewUserService(userRepository)

	// Prefijo comun a mis endpoints de user
	api_user := app.Group("/api/user")

	// // Prefijo comun a mis endpoints de product
	// api_product := app.Group("/api/product")

	// // Prefijo comun a mis endpoints de order
	// api_order := app.Group("/api/order")

	//Aca importo todos los endpoints de user
	routes.UserRoutes(api_user, userService)

	// //Aca importo todos los endpoints de product
	// routes.ProductRoutes(api_product)

	// //Aca importo todos los endpoints de order
	// routes.OrderRoutes(api_order)

	return app

}

func main() {
	clienteSQL := database.ConnectDb()

	app := InicializarApp(http.DefaultClient, clienteSQL)

	log.Fatal(app.Listen(":3000"))
}

type internalError struct {
	Message string `json:"message"`
}
