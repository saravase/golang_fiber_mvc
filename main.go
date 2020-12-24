package main

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/go-playground/validator.v9"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbClient *gorm.DB
)

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name" validate:"required,min=3,max=32"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func init() {
	var err error
	dsn := "host=localhost user=primz password=primz@2207 dbname=primz port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	dbClient, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func main() {
	app := fiber.New()

	app.Get("/user/:id", handlerGetUser)
	app.Post("/user", handlerCreateUser)

	app.Listen(":9090")
}

func handlerGetUser(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return handleError(c, fiber.StatusBadRequest, "Invalid user id")
	}

	user, err := GetUser(id)
	if err != nil {
		return handleError(c, fiber.StatusNotFound, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": user,
	})
}

func ValidateStruct(data interface{}) []*ErrorResponse {

	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func handlerCreateUser(c *fiber.Ctx) error {

	user := new(User)
	if err := c.BodyParser(user); err == nil {
		errors := ValidateStruct(user)
		if errors != nil {
			return handleError(c, fiber.StatusBadRequest, errors)
		}
	}

	if err := SaveUser(user); err != nil {
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

func SaveUser(user *User) error {

	if user == nil {
		return errors.New("Invalid user to save")
	}

	result := dbClient.Create(user)
	if err := result.Error; err != nil {
		return err
	}

	return nil
}

func GetUser(id int64) (*User, error) {

	var user = new(User)
	result := dbClient.Where("id = ?", id).Find(&user)
	if err := result.Error; err != nil {
		return nil, err
	}

	if user.Id == 0 {
		return nil, errors.New("User not found")
	}

	return user, nil
}
