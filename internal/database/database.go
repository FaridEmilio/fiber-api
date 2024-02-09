package internal

import (
	"log"
	"os"

	pkg "github.com/faridEmilio/fiber-api/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("No se pudo conectar a la base de datos", err.Error())
		os.Exit(2)
	}

	log.Println("Conexión exitosa a la base de datos")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")

	db.AutoMigrate(&pkg.User{}, &pkg.Product{}, &pkg.Order{})

	Database = DbInstance{Db: db}
}
