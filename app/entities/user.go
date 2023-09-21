package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// UserEntity is Unique app_id & phone_number or app_id & email
type UserEntity struct {
	gorm.Model
	ID                         uuid.UUID `gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	AppID                      uuid.UUID `gorm:"uniqueIndex:idx_app_phone;uniqueIndex:idx_app_email"`
	App                        AppEntity `gorm:"foreignKey:AppID"`
	DisplayName                string
	FirstName                  string
	LastName                   string
	PhoneNumber                string `gorm:"uniqueIndex:idx_app_phone"`
	Email                      string `gorm:"uniqueIndex:idx_app_email"`
	Password                   string
	Ticket                     string
	TicketExpired              time.Time
	VerifyFlag                 bool `gorm:"default:false"`
	RequiredChangePasswordFlag bool `gorm:"default:false"`
	MetaData                   datatypes.JSON
}
