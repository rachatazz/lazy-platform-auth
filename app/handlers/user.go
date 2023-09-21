package handler

import (
	"strings"

	model "lazy-platform-auth/app/models"
	service "lazy-platform-auth/app/services"
	"lazy-platform-auth/utils"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userService service.UserService
}

type UserHandler interface {
	CreateUser(*fiber.Ctx) error
	GetUser(*fiber.Ctx) error
	Token(*fiber.Ctx) error
	UpdateUser(*fiber.Ctx) error
	ChangePassword(*fiber.Ctx) error
	ForgotPassword(*fiber.Ctx) error
	ResetPassword(*fiber.Ctx) error
	RefreshToken(*fiber.Ctx) error
	DeleteToken(*fiber.Ctx) error
	GetMe(*fiber.Ctx) error
	GetTemplate(*fiber.Ctx) error
}

func NewUserHandler(userService service.UserService) UserHandler {
	return userHandler{userService: userService}
}

func (h userHandler) CreateUser(c *fiber.Ctx) error {
	appId := c.Get("app-id")
	body := model.CreateUserRequest{}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.ErrUnprocessableEntity.Code).JSON(fiber.Map{
			"status":  "error",
			"message": "body param is invalid",
		})
	}
	body.Email = strings.TrimSpace(strings.ToLower(body.Email))

	user, err := h.userService.CreateUser(appId, body)
	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status":  "error",
			"message": err.Message,
		})
	}
	return c.JSON(user)
}

func (h userHandler) GetMe(c *fiber.Ctx) error {
	id := utils.CliamsToken(c)
	user, err := h.userService.GetUserById(id)
	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status":  "error",
			"message": err.Message,
		})
	}

	return c.JSON(user)
}

func (h userHandler) GetUser(c *fiber.Ctx) error {
	query := model.UserQuery{}

	if err := c.QueryParser(&query); err != nil {
		return c.Status(fiber.ErrUnprocessableEntity.Code).JSON(fiber.Map{
			"status":  "error",
			"message": "query param is invalid",
		})
	}

	user, err := h.userService.GetUser(query)
	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status":  "error",
			"message": err.Message,
		})
	}

	return c.JSON(user)
}

func (h userHandler) Token(c *fiber.Ctx) error {
	appId := c.Get("app-id")
	body := model.TokenRequest{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.ErrUnprocessableEntity.Code).JSON(fiber.Map{
			"status":  "error",
			"message": "body param is invalid",
		})
	}
	body.Email = strings.TrimSpace(strings.ToLower(body.Email))

	session, err := h.userService.Token(appId, body)
	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status":  "error",
			"message": err.Message,
		})
	}

	return c.JSON(&session)
}

func (h userHandler) UpdateUser(c *fiber.Ctx) error {
	id := utils.CliamsToken(c)
	body := model.UpdateUserRequest{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.ErrUnprocessableEntity.Code).JSON(fiber.Map{
			"status":  "error",
			"message": "body param is invalid",
		})
	}

	user, err := h.userService.UpdateUser(id, body)
	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status":  "error",
			"message": err.Message,
		})
	}

	return c.JSON(user)
}

func (h userHandler) RefreshToken(c *fiber.Ctx) error {
	body := model.RefreshTokenRequest{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.ErrUnprocessableEntity.Code).JSON(fiber.Map{
			"status":  "error",
			"message": "body param is invalid",
		})
	}

	session, err := h.userService.RefreshToken(body)
	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status":  "error",
			"message": err.Message,
		})
	}

	return c.JSON(&session)
}

func (h userHandler) DeleteToken(c *fiber.Ctx) error {
	id := utils.CliamsToken(c)
	err := h.userService.DeleteToken(id)
	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status":  "error",
			"message": err.Message,
		})
	}

	return c.JSON(fiber.Map{"status": "ok"})
}

func (h userHandler) ChangePassword(c *fiber.Ctx) error {
	id := utils.CliamsToken(c)
	body := model.ChangePasswordRequest{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.ErrUnprocessableEntity.Code).JSON(fiber.Map{
			"status":  "error",
			"message": "body param is invalid",
		})
	}

	err := h.userService.ChangePassword(id, body)
	if err != nil {
		return c.Status(fiber.ErrUnprocessableEntity.Code).JSON(fiber.Map{
			"status":  "error",
			"message": err.Message,
		})
	}

	return c.JSON(fiber.Map{"status": "ok"})
}

func (h userHandler) ForgotPassword(c *fiber.Ctx) error {
	appId := c.Get("app-id")
	body := model.ForgotPasswordRequest{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.ErrUnprocessableEntity.Code).JSON(fiber.Map{
			"status":  "error",
			"message": "body param is invalid",
		})
	}
	body.Email = strings.TrimSpace(strings.ToLower(body.Email))

	err := h.userService.ForgotPassword(appId, body.Email)
	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status":  "error",
			"message": err.Message,
		})
	}
	return c.JSON(fiber.Map{"status": "ok"})
}

func (h userHandler) ResetPassword(c *fiber.Ctx) error {
	body := model.ResetPasswordRequest{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.ErrUnprocessableEntity.Code).JSON(fiber.Map{
			"status":  "error",
			"message": "body param is invalid",
		})
	}

	err := h.userService.ResetPassword(body)
	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status":  "error",
			"message": err.Message,
		})
	}
	return c.JSON(fiber.Map{"status": "ok"})
}

func (h userHandler) GetTemplate(c *fiber.Ctx) error {
	res, err := h.userService.GetTemplate()
	if err != nil {
		return c.Status(fiber.ErrUnprocessableEntity.Code).JSON(fiber.Map{
			"status":  "error",
			"message": err.Message,
		})
	}
	return c.JSON(&res)
}
