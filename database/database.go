package database

import (
	"jsonapi/model"

	"github.com/jinzhu/gorm"
)

// Global database connection
var DB gorm.DB

// AutoMigrate creates or modifies existing tables
func AutoMigrate(DBObject gorm.DB) {
	DBObject.AutoMigrate(
		&model.User{},
	)
}
