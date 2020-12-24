package controller

import (
	"log"
	"strconv"

	"github.com/saravase/golang_fiber_mvc/mvc/services"

	"github.com/saravase/golang_fiber_mvc/mvc/domain"

	fiber "github.com/gofiber/fiber/v2"
)

var (
	UserController = userController{}
)

type userController struct{}

func (controller userController) Get(c *fiber.Ctx) error {

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return handleError(c, fiber.StatusBadRequest, "Invalid user id")
	}

	user, err := services.UserService.Get(id)
	if err != nil {
		return handleError(c, fiber.StatusNotFound, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": user,
	})
}

func (controller userController) Create(c *fiber.Ctx) error {

	user := new(domain.User)
	if err := c.BodyParser(user); err == nil {
		errors := ValidateStruct(user)
		if errors != nil {
			return handleError(c, fiber.StatusBadRequest, errors)
		}
	}

	log.Printf("User: %v\n", user)

	if err := services.UserService.Save(user); err != nil {
		return handleError(c, fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": user.Id,
	})
}

func handleError(c *fiber.Ctx, status int, errRes interface{}) error {
	return c.Status(status).JSON(fiber.Map{
		"message": errRes,
	})
}
