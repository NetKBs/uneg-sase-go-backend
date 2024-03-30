package university

type Career struct {
	Id   uint `gorm:"primaryKey"`
	Name string
}
