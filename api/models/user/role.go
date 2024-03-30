package user

type Role struct {
	Id   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique"`
}

/*
	Inputs Models
*/

type RoleInputCreate struct {
	Name string `json:"name" binding:"required"`
}

type RoleInputUpdate struct {
	Id   uint   `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type RoleInputDelete struct {
	Id uint `json:"id" binding:"required"`
}
