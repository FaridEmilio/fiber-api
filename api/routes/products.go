package routes

import (
	"errors"

	internal "github.com/faridEmilio/fiber-api/internal/database"
	pkg "github.com/faridEmilio/fiber-api/pkg/models"
	"github.com/gofiber/fiber/v2"
)

type Product struct {
	ID           uint   `json: "id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

// Endpoints de product
func ProductRoutes(app fiber.Router) {
	app.Post("/new", CreateProduct())
	app.Get("/find/:id", GetProductByID())
	app.Get("/list", GetProducts())
	app.Put("/update/:id", UpdateProduct())
	app.Delete("/delete/:id", DeleteProduct())
}

func CreateResponseProduct(productModel pkg.Product) Product {
	return Product{ID: productModel.ID, Name: productModel.Name, SerialNumber: productModel.SerialNumber}
}

func CreateProduct() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var product pkg.Product

		if err := ctx.BodyParser(&product); err != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		internal.Database.Db.Create(&product)

		responseProduct := CreateResponseProduct(product)

		return ctx.Status(200).JSON(responseProduct)
	}
}

// No olvidar el puntero
func findProduct(id int, product *pkg.Product) error {

	internal.Database.Db.Find(&product, "id= ?", id)

	if product.ID == 0 {
		return errors.New("Product does not exist")
	}
	return nil
}

func GetProductByID() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, err := ctx.ParamsInt("id")

		var product pkg.Product

		if err != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := findProduct(id, &product); err != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		responseProduct := CreateResponseProduct(product)

		return ctx.Status(200).JSON(responseProduct)
	}
}

func GetProducts() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		products := []pkg.Product{}

		internal.Database.Db.Find(&products)

		responseProducts := []Product{}

		for _, product := range products {
			responseProduct := CreateResponseProduct(product)
			responseProducts = append(responseProducts, responseProduct)
		}

		return ctx.Status(200).JSON(responseProducts)
	}
}

func UpdateProduct() fiber.Handler {

	return func(ctx *fiber.Ctx) error {
		id, err := ctx.ParamsInt("id")

		var product pkg.Product

		if err != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := findProduct(id, &product); err != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		type updateProduct struct {
			Name         string `json:"name"`
			SerialNumber string `json:"serial_number"`
		}

		var updateData updateProduct

		if err := ctx.BodyParser(&updateData); err != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		product.Name = updateData.Name
		product.SerialNumber = updateData.SerialNumber

		internal.Database.Db.Save(&product)

		responseProduct := CreateResponseProduct(product)

		return ctx.Status(200).JSON(responseProduct)
	}

}

func DeleteProduct() fiber.Handler {

	return func(ctx *fiber.Ctx) error {
		id, err := ctx.ParamsInt("id")

		var product pkg.Product

		if err != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := findProduct(id, &product); err != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := internal.Database.Db.Delete(&product).Error; err != nil {
			return ctx.Status(404).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return ctx.Status(200).SendString("Producto eliminado con éxito")

	}
}
