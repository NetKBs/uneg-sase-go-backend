package initializers

import (
	uni "github.com/NetKBs/uneg-sase-go-backend/api/models/university"
	"github.com/NetKBs/uneg-sase-go-backend/api/models/user"
)

func SyncDatabase() {

	models := []interface{}{
		&user.User{},
		&user.Role{},
		&user.Student{},
		&user.Teacher{},
		&user.AcademicData{},
		&user.Gender{},
		&user.CivilState{},
		&user.PersonalData{},

		&uni.Campus{},
		&uni.Career{},
	}

	for _, model := range models {
		DB.AutoMigrate(model)
	}

}
