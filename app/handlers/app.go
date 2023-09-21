package handler

import (
	model "lazy-platform-auth/app/models"
	service "lazy-platform-auth/app/services"

	"github.com/gofiber/fiber/v2"
)

type appHandler struct {
	appService service.AppService
}

type AppHandler interface {
	CreateApp(c *fiber.Ctx) error
	UpdateApp(c *fiber.Ctx) error
	DeleteApp(c *fiber.Ctx) error
	GetApp(c *fiber.Ctx) error
	GetAppById(c *fiber.Ctx) error
}

func NewAppHandler(appService service.AppService) AppHandler {
	return appHandler{appService: appService}
}

func (h appHandler) CreateApp(c *fiber.Ctx) error {
	body := model.CreateAppRequest{}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.ErrUnprocessableEntity.Code).JSON(fiber.Map{
			"status":  "error",
			"message": "body param is invalid",
		})
	}

	app, err := h.appService.CreateApp(body.Name)
	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status":  "error",
			"message": err.Message,
		})
	}

	return c.JSON(app)
}

func (h appHandler) UpdateApp(c *fiber.Ctx) error {
	id := c.Params("id")
	body := model.UpdateAppRequest{}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.ErrUnprocessableEntity.Code).JSON(fiber.Map{
			"status":  "error",
			"message": "body param is invalid",
		})
	}

	app, err := h.appService.UpdateApp(id, body.Name)
	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status":  "error",
			"message": err.Message,
		})
	}

	return c.JSON(app)
}

func (h appHandler) DeleteApp(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.appService.DeleteApp(id)
	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status":  "error",
			"message": err.Message,
		})
	}
	return c.JSON(fiber.Map{"status": "ok"})
}

func (h appHandler) GetApp(c *fiber.Ctx) error {
	apps, err := h.appService.GetApps()
	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status":  "error",
			"message": err.Message,
		})
	}
	return c.JSON(apps)
}

func (h appHandler) GetAppById(c *fiber.Ctx) error {
	id := c.Params("id")
	app, err := h.appService.GetAppById(id)
	if err != nil {
		return c.Status(err.Code).JSON(fiber.Map{
			"status":  "error",
			"message": err.Message,
		})
	}
	return c.JSON(app)
}
