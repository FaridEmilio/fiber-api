package routes

import (
	"errors"

	internal "github.com/faridEmilio/fiber-api/internal/database"
	pkg "github.com/faridEmilio/fiber-api/pkg/models"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	//este no es el model User, se lo puede ver como el serializable
	ID        uint   `json: "id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// Aca van todos los endpoints de /user
func UserRoutes(app fiber.Router) {
	app.Post("/new", CreateUser())
	app.Get("/list", GetUsers())
	app.Get("/find/:id", GetUserById())
	app.Put("/update/:id", UpdateUser())
	app.Delete("/delete/:id", DeleteUSer())
}

func CreateResponseUser(userModel pkg.User) User {
	return User{ID: userModel.ID, FirstName: userModel.FirstName, LastName: userModel.LastName}
}

/*
La función CreateUser() ahora devuelve una función que coincide con la firma de fiber.Handler.
Dentro de esta función anónima, se realiza la lógica de manejo de solicitudes HTTP.
Si hay un error al analizar el cuerpo de la solicitud, se responde con un código de estado 400 y el mensaje de error.
Después de crear el usuario en la base de datos y preparar la respuesta del usuario, se responde con un código de estado 200
y los datos del usuario en formato JSON.
*/

func CreateUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var user pkg.User

		if err := ctx.BodyParser(&user); err != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		internal.Database.Db.Create(&user)
		responseUser := CreateResponseUser(user)

		return ctx.Status(200).JSON(responseUser)
	}
}

func findUser(id int, user *pkg.User) error {
	internal.Database.Db.Find(&user, "id= ?", id)

	if user.ID == 0 {
		return errors.New("User does not exist")
	}

	return nil
}

func GetUserById() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, err := ctx.ParamsInt("id")

		var user pkg.User

		if err != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := findUser(id, &user); err != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		responseUser := CreateResponseUser(user)

		return ctx.Status(200).JSON(responseUser)
	}
}

func GetUsers() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		users := []pkg.User{}
		internal.Database.Db.Find(&users)

		responseUsers := []User{}

		for _, user := range users {
			responseUser := CreateResponseUser(user)
			responseUsers = append(responseUsers, responseUser)
		}

		return ctx.Status(200).JSON(responseUsers)
	}
}

func UpdateUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, err := ctx.ParamsInt("id")

		var user pkg.User

		if err != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := findUser(id, &user); err != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		type UpdateUser struct {
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
		}

		var updateData UpdateUser

		if err := ctx.BodyParser(&updateData); err != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		user.FirstName = updateData.FirstName
		user.LastName = updateData.LastName

		internal.Database.Db.Save(&user)

		responseUser := CreateResponseUser(user)

		return ctx.Status(200).JSON(responseUser)

	}
}

func DeleteUSer() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, err := ctx.ParamsInt("id")

		var user pkg.User

		if err != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := findUser(id, &user); err != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := internal.Database.Db.Delete(&user).Error; err != nil {
			return ctx.Status(404).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return ctx.Status(200).SendString("Usuario eliminado con éxito")

	}
}
