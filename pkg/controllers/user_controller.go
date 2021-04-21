package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wtfcode/hermes-api/pkg/models"
	"github.com/wtfcode/hermes-api/pkg/services"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type IUserController interface {
	GetAll() fiber.Handler
	GetOne() fiber.Handler
	Save() fiber.Handler
	Delete() fiber.Handler
	Update() fiber.Handler
}

type UserController struct {
	s services.IUserService
}

func NewUserController(db *gorm.DB) IUserController {
	return &UserController{
		s: services.NewUserService(db),
	}
}

func (u *UserController) GetAll() fiber.Handler {
	return func(c *fiber.Ctx) error {
		page, err := c.ParamsInt("page")
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		users := u.s.GetAllUsers(page)
		return c.JSON(users)
	}
}

func (u *UserController) GetOne() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		user, err := u.s.GetByID(id)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.JSON(user)
	}
}

func (u *UserController) Save() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user models.User
		if err := c.BodyParser(&user); err != nil {
			return err
		}
		pwd, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
		user.Password = string(pwd)
		if err := u.s.CreateUser(user); err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.SendStatus(fiber.StatusCreated)
	}
}

func (u *UserController) Delete() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		if err := u.s.DeleteUser(id); err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.SendStatus(fiber.StatusOK)
	}
}

func (u *UserController) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user models.User
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		if err := c.BodyParser(&user); err != nil {
			return c.SendStatus(fiber.StatusUnprocessableEntity)
		}
		if err := u.s.UpdateUser(user, uint(id)); err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.SendStatus(fiber.StatusOK)
	}
}
