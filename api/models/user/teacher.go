package user

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	User   User
	UserID uint
}
