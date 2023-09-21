package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RefreshTokenEntity struct {
	gorm.Model
	ID           uuid.UUID `gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	UserID       uuid.UUID
	RefreshToken string
	ExpiresAt    time.Time
}
