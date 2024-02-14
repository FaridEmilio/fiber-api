package internal

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/faridEmilio/fiber-api/internal/config"
	pkg "github.com/faridEmilio/fiber-api/pkg/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Instancia de base de datos
type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", config.USER, config.PASSW, config.HOST, config.PORT, config.DB)
	// logs.Info(dsn)
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			loc, _ := time.LoadLocation("America/Argentina/Buenos_Aires")
			return time.Now().In(loc)
		},
	})

	//db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("No se pudo conectar a la base de datos", err.Error())
		os.Exit(2)
	}

	log.Println("Conexión exitosa a la base de datos")
	//db.Logger = logger.Default.LogMode(logger.Info)

	//Crea automáticamente tablas en la base de datos correspondientes a las estructuras de datos definidas
	log.Println("Running Migrations")
	gormDB.AutoMigrate(&pkg.User{}, &pkg.Product{}, &pkg.Order{})

	Database = DbInstance{Db: gormDB}
}
