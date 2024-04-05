package user

type Role struct {
	Id   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"unique" json:"name"`
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
