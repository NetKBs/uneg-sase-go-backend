package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	RoleId   uint
	Role     Role
}

type UserInputCreate struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	RoleId   uint   `json:"role_id" binding:"required"`
}

type UserInputUpdate struct {
	Id       uint   `json:"id" binding:"required"`
	Username string `json:"username"`
	Password string `json:"password"`
	RoleId   uint   `json:"role_id"`
}

type UserInputDelete struct {
	Id uint `json:"id" binding:"required"`
}
