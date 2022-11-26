package controller

import (
	"github.com/gofiber/fiber/v2"
)

// UserInterfaceController is an interface for user controller
// Path: interface\controller\user_controller.go
// Used to create logic and throw to router

type UserInterfaceController interface {
	GetUsers(context *fiber.Ctx) error
	CreateUser(context *fiber.Ctx) error
	GetUserById(context *fiber.Ctx) error
    UpdateUser(context *fiber.Ctx) error
    DeleteById(context *fiber.Ctx) error
}
