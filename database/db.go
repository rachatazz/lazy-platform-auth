package database

import (
	"lazy-platform-auth/config"
	"lazy-platform-auth/logs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(configService config.Config) *gorm.DB {
	dsn := configService.DataBaseUri
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{DryRun: false, Logger: logs.DbLogger()})
	if err != nil {
		panic(err)
	}

	return db
}
