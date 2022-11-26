package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
	"tutorial/infrastructure/datastore"
    "tutorial/infrastructure/migration"
    "tutorial/infrastructure/router"
	"tutorial/registry"
)

func main() {

	customLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Microsecond,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)
	db := datastore.Database()
    migration.Run()
	db.Logger = customLogger

	app := fiber.New()
	r := registry.BaseRegistry(db)

	router.Router(app, r.RegistryAppController())

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
