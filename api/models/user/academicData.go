package user

import "github.com/NetKBs/uneg-sase-go-backend/api/models/university"

type AcademicData struct {
	Id        uint `gorm:"primaryKey"`
	Active    bool
	UnegEmail string
	Campus    university.Campus
	CampusId  uint
	Career    university.Career
	CareerId  uint
}
