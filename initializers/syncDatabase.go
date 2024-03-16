package initializers

import "github.com/NetKBs/uneg-sase-go-backend/api/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
