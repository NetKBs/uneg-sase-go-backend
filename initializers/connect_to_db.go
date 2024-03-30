package initializers

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	DB, err = gorm.Open(sqlite.Open("uneg_db.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}
