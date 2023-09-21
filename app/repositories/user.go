package repository

import (
	"strings"

	entity "lazy-platform-auth/app/entities"
	model "lazy-platform-auth/app/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(entity.UserEntity) (*entity.UserEntity, error)
	Update(entity.UserEntity) (*entity.UserEntity, error)
	Find(model.UserQuery) ([]entity.UserEntity, error)
	FindByEmailOrPhoneNumber(
		appId string,
		email string,
		phoneNumber string,
	) (*entity.UserEntity, error)
	FindByEmail(appId string, email string) (*entity.UserEntity, error)
	FindById(string) (*entity.UserEntity, error)
	FindByTicket(ticket string) (*entity.UserEntity, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{db: db}
}

func (r userRepository) Create(user entity.UserEntity) (*entity.UserEntity, error) {
	tx := r.db.Create(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

func (r userRepository) Update(user entity.UserEntity) (*entity.UserEntity, error) {
	tx := r.db.Save(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}

func (r userRepository) FindById(id string) (*entity.UserEntity, error) {
	user := entity.UserEntity{}
	tx := r.db.Where("id = ?", id).Take(&user)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

func (r userRepository) Find(query model.UserQuery) ([]entity.UserEntity, error) {
	users := []entity.UserEntity{}
	tx := r.db

	if query.Limit != 0 {
		tx = tx.Limit(query.Limit)
	}

	if query.Offser != 0 {
		tx = tx.Offset(query.Offser)
	}

	if query.OrderBy != "" {
		orderStr := query.OrderBy + " "
		if query.SortBy != "" {
			orderStr = orderStr + strings.ToUpper(query.SortBy)
		} else {
			orderStr = orderStr + "DESC"
		}
		tx = tx.Order(orderStr)
	}

	if query.ID != "" {
		tx = tx.Where("id = ?", query.ID)
	}

	if query.AppID != "" {
		tx = tx.Where("app_id = ?", query.AppID)
	}

	if query.Keyword != "" {
		tx = tx.Where(
			"display_name ? OR first_name ILIKE ? OR last_name ILIKE ? OR phone_number ILIKE ?",
			"%"+query.Keyword+"%",
			"%"+query.Keyword+"%",
			"%"+query.Keyword+"%",
			"%"+query.Keyword+"%",
		)
	}

	tx.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil
}

func (r userRepository) FindByEmailOrPhoneNumber(
	appId string,
	email string,
	phoneNumber string,
) (*entity.UserEntity, error) {
	user := entity.UserEntity{}

	tx := r.db.Where("app_id = ? AND email = ?", appId, email).
		Or("app_id = ? AND phone_number = ?", appId, phoneNumber).
		Take(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}

func (r userRepository) FindByEmail(appId string, email string) (*entity.UserEntity, error) {
	user := entity.UserEntity{}
	tx := r.db.Where("app_id = ? AND email = ?", appId, email).Take(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

func (r userRepository) FindByTicket(ticket string) (*entity.UserEntity, error) {
	user := entity.UserEntity{}
	tx := r.db.Where("ticket = ?", ticket).Take(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}
