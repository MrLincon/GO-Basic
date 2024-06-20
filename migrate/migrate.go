package migrate

import (
	"GO-Basic/initializers"
	"GO-Basic/models"
)

func Migrate() {
	err := initializers.DB.AutoMigrate(&models.Post{})
	if err != nil {
		return
	}
}
