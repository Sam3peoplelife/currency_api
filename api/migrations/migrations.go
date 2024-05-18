package migrations

import (
    "api/models"
    "gorm.io/gorm"
)

//DB migration
func Migrate(db *gorm.DB) {
    db.AutoMigrate(&models.User{})
}

