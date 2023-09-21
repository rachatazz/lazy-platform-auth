package service

import (
	entity "lazy-platform-auth/app/entities"
	model "lazy-platform-auth/app/models"
	repository "lazy-platform-auth/app/repositories"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AppService interface {
	CreateApp(name string) (*model.AppResponse, *fiber.Error)
	UpdateApp(id string, name string) (*model.AppResponse, *fiber.Error)
	DeleteApp(id string) *fiber.Error
	GetApps() ([]model.AppResponse, *fiber.Error)
	GetAppById(id string) (*model.AppResponse, *fiber.Error)
}

type appService struct {
	appRepository repository.AppRepository
}

func NewAppService(appRepository repository.AppRepository) AppService {
	return appService{appRepository: appRepository}
}

func (s appService) CreateApp(name string) (*model.AppResponse, *fiber.Error) {
	app, err := s.appRepository.Create(entity.AppEntity{Name: name})
	if err != nil {
		return nil, fiber.NewError(
			fiber.ErrUnprocessableEntity.Code,
			"something went wrong",
		)
	}

	res := model.AppResponse{
		ID:        app.ID.String(),
		Name:      app.Name,
		CreatedAt: app.CreatedAt.String(),
		UpdatedAt: app.UpdatedAt.String(),
	}
	return &res, nil
}

func (s appService) UpdateApp(id string, name string) (*model.AppResponse, *fiber.Error) {
	appId, err := uuid.Parse(id)
	if err != nil {
		return nil, fiber.NewError(fiber.ErrInternalServerError.Code, "UUID is incorrect")
	}

	app, err := s.appRepository.Update(entity.AppEntity{ID: appId, Name: name})
	if err != nil {
		return nil, fiber.NewError(
			fiber.ErrUnprocessableEntity.Code,
			"something went wrong",
		)
	}

	res := model.AppResponse{
		ID:        app.ID.String(),
		Name:      app.Name,
		CreatedAt: app.CreatedAt.String(),
		UpdatedAt: app.UpdatedAt.String(),
	}

	return &res, nil
}

func (s appService) DeleteApp(id string) *fiber.Error {
	_, err := s.appRepository.Delete(id)
	if err != nil {
		return fiber.NewError(fiber.ErrUnauthorized.Code, "app not found")
	}
	return nil
}

func (s appService) GetApps() ([]model.AppResponse, *fiber.Error) {
	apps, err := s.appRepository.Find()
	if err != nil {
		return nil, fiber.NewError(
			fiber.ErrUnprocessableEntity.Code,
			"something went wrong",
		)
	}

	res := []model.AppResponse{}
	for _, app := range apps {
		res = append(res, model.AppResponse{
			ID:        app.ID.String(),
			Name:      app.Name,
			CreatedAt: app.CreatedAt.String(),
			UpdatedAt: app.UpdatedAt.String(),
		})
	}
	return res, nil
}

func (s appService) GetAppById(id string) (*model.AppResponse, *fiber.Error) {
	app, err := s.appRepository.FindById(id)
	if err != nil {
		return nil, fiber.NewError(
			fiber.ErrNotFound.Code,
			"app not found",
		)
	}
	res := model.AppResponse{
		ID:        app.ID.String(),
		Name:      app.Name,
		CreatedAt: app.CreatedAt.String(),
		UpdatedAt: app.UpdatedAt.String(),
	}
	return &res, nil
}
