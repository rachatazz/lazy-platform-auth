package repository

import (
	entity "lazy-platform-auth/app/entities"

	"gorm.io/gorm"
)

type appRepository struct {
	db *gorm.DB
}

type AppRepository interface {
	Create(app entity.AppEntity) (*entity.AppEntity, error)
	Update(app entity.AppEntity) (*entity.AppEntity, error)
	Delete(id string) (*entity.AppEntity, error)
	Find() ([]entity.AppEntity, error)
	FindById(id string) (*entity.AppEntity, error)
}

func NewAppRepository(db *gorm.DB) AppRepository {
	return appRepository{db: db}
}

func (r appRepository) Create(app entity.AppEntity) (*entity.AppEntity, error) {
	tx := r.db.Create(&app)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &app, nil
}

func (r appRepository) Update(app entity.AppEntity) (*entity.AppEntity, error) {
	tx := r.db.Save(&app)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &app, nil
}

func (r appRepository) Delete(id string) (*entity.AppEntity, error) {
	app := entity.AppEntity{}

	tx := r.db.Delete(&app, "id = ?", id)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &app, nil
}

func (r appRepository) Find() ([]entity.AppEntity, error) {
	apps := []entity.AppEntity{}
	tx := r.db
	tx.Find(&apps)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return apps, nil
}

func (r appRepository) FindById(id string) (*entity.AppEntity, error) {
	app := entity.AppEntity{}
	tx := r.db.Take(&app, "id = ?", id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &app, nil
}
