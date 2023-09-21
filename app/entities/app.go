package entity

import (
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type AppEntity struct {
	gorm.Model
	ID   uuid.UUID `gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	Name string
}
