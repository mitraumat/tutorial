package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"tutorial/infrastructure/helper"
)

func Router(route *fiber.App, c helper.Application) *fiber.App {
    route.Use(logger.New())
    route.Use(recover.New())

    route.Get("/users", c.User.GetUsers)
    route.Get("/users/:id", c.User.GetUserById)
    route.Put("/users/:id",c.User.UpdateUser)
    route.Post("/users/create", c.User.CreateUser)
    route.Delete("users/:id",c.User.DeleteById)


    return route
}
