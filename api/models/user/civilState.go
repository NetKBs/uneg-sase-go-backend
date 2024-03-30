package user

type CivilState struct {
	Id   uint `gorm:"primaryKey"`
	Name string
}
