package repository

import (
	entity "lazy-platform-auth/app/entities"

	"gorm.io/gorm"
)

type sessionRepository struct {
	db *gorm.DB
}

type SessionRepository interface {
	Create(session entity.RefreshTokenEntity) (*entity.RefreshTokenEntity, error)
	Update(session entity.RefreshTokenEntity) (*entity.RefreshTokenEntity, error)
	Delete(id string) (*entity.RefreshTokenEntity, error)
	FindByAccountId(id string) (*entity.RefreshTokenEntity, error)
	Upsert(session entity.RefreshTokenEntity) (*entity.RefreshTokenEntity, error)
}

func NewSessionRepository(db *gorm.DB) SessionRepository {
	return sessionRepository{db: db}
}

func (r sessionRepository) Create(
	session entity.RefreshTokenEntity,
) (*entity.RefreshTokenEntity, error) {
	tx := r.db.Create(&session)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &session, nil
}

func (r sessionRepository) Update(
	session entity.RefreshTokenEntity,
) (*entity.RefreshTokenEntity, error) {
	return nil, nil
}

func (r sessionRepository) Delete(id string) (*entity.RefreshTokenEntity, error) {
	session := entity.RefreshTokenEntity{}

	tx := r.db.Where("user_id = ?", id).Delete(&session)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &session, nil
}

func (r sessionRepository) FindByAccountId(id string) (*entity.RefreshTokenEntity, error) {
	session := entity.RefreshTokenEntity{}

	tx := r.db.Where("user_id = ?", id).Take(&session)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &session, nil
}

func (r sessionRepository) Upsert(
	session entity.RefreshTokenEntity,
) (*entity.RefreshTokenEntity, error) {
	r.db.Where("user_id = ?", session.UserID).Take(&session)

	tx := r.db.Save(&session)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &session, nil
}
