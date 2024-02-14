package routes

import (
	"errors"

	internal "github.com/faridEmilio/fiber-api/internal/database"
	pkg "github.com/faridEmilio/fiber-api/pkg/entities"
	"github.com/gofiber/fiber/v2"
)

type Order struct {
	ID      uint    `json: "id"`
	User    User    `json:"user"`
	Product Product `json:"product"`
}

// EndPoints de order
func OrderRoutes(app fiber.Router) {
	app.Post("/new", CreateOrder())
	app.Get("/find/:id", GetOrderByID())
	app.Get("/list", GetOrders())

}

func CreateResponseOrder(orderModel pkg.Order, user User, product Product) Order {
	return Order{ID: orderModel.ID, User: user, Product: product}
}
func CreateOrder() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var order pkg.Order

		if err := ctx.BodyParser(&order); err != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		// En este caso debemos agregar validaciones
		// Es decir, que para poder crear una orden debe exister el producto y el usuario
		var user pkg.User
		if err := findUser(order.UserRefer, &user); err != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		var product pkg.Product
		if err := findProduct(order.ProductRefer, &product); err != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		//Una vez validado todo, se guarda
		internal.Database.Db.Save(&order)

		responseUser := CreateResponseUser(user)
		responseProduct := CreateResponseProduct(product)
		responseOrder := CreateResponseOrder(order, responseUser, responseProduct)

		return ctx.Status(200).JSON(responseOrder)
	}
}

func findOrder(id int, order *pkg.Order) error {
	internal.Database.Db.Find(&order, "id= ?", id)

	if order.ID == 0 {
		return errors.New("Order does not exist")
	}

	return nil
}

func GetOrderByID() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, err := ctx.ParamsInt("id")

		var order pkg.Order

		if err != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := findOrder(id, &order); err != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		var user pkg.User
		var product pkg.Product

		internal.Database.Db.First(&user, order.UserRefer)
		internal.Database.Db.First(&product, order.ProductRefer)

		responseUser := CreateResponseUser(user)
		responseProduct := CreateResponseProduct(product)

		responseOrder := CreateResponseOrder(order, responseUser, responseProduct)

		return ctx.Status(200).JSON(responseOrder)
	}
}

func GetOrders() fiber.Handler {

	return func(ctx *fiber.Ctx) error {

		orders := []pkg.Order{}

		internal.Database.Db.Find(&orders)

		responseOrders := []Order{}

		for _, order := range orders {
			var user pkg.User
			var product pkg.Product

			internal.Database.Db.Find(&user, "id = ?", order.UserRefer)
			internal.Database.Db.Find(&product, "id = ?", order.ProductRefer)

			responseUser := CreateResponseUser(user)
			responseProduct := CreateResponseProduct(product)
			responseOrder := CreateResponseOrder(order, responseUser, responseProduct)

			responseOrders = append(responseOrders, responseOrder)

		}

		return ctx.Status(200).JSON(responseOrders)
	}
}
