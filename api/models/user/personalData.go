package user

type PersonalData struct {
	Id             uint `gorm:"primaryKey"`
	Name           string
	SecondName     string
	LastName       string
	SecondLastName string
	Gender         Gender
	GenderId       uint
	BornDate       string `gorm:"type:date"`
	Ic             uint   `gorm:"not null;unique"`
	CivilState     CivilState
	CivilStateId   uint
	Email          string `gorm:"not null"`
	AltEmail       string
	Phone          uint `gorm:"not null"`
	AltPhone       uint
	BorthCity      string
	BorthState     string
	BorthCountry   string
	HouseDir       string
	JobDir         string
}
