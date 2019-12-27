package migrate

import (
	"webrtc-server/internal/models"

	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	db.Debug().AutoMigrate(&models.Error{}, &models.Todo{})
}
