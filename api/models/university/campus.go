package university

type Campus struct {
	Id   uint `gorm:"primaryKey"`
	Name string
}
