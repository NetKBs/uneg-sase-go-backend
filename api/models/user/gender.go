package user

type Gender struct {
	Id   uint `gorm:"primaryKey"`
	Name string
}
