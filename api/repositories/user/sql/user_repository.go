package sql

import (
	"errors"

	"github.com/NetKBs/uneg-sase-go-backend/api/models/user"
	"github.com/NetKBs/uneg-sase-go-backend/initializers"
)

type UserRepository struct {
}

func (r UserRepository) Create(u user.User) (user.User, error) {

	result := initializers.DB.Create(&u)

	if result.Error != nil {
		return user.User{}, result.Error
	}

	return u, nil
}

func (r UserRepository) Get(id uint) (user.User, error) {
	var u user.User
	result := initializers.DB.First(&u, id)

	if result.Error != nil {
		return user.User{}, result.Error
	}

	return u, nil
}

func (r UserRepository) GetAll() ([]user.User, error) {

	var users []user.User
	result := initializers.DB.Find(&users)

	if result.Error != nil {
		return []user.User{}, result.Error
	}

	return users, nil
}

func (r UserRepository) Update(u user.User) (user.User, error) {

	result := initializers.DB.Model(&u).Where("Id = ?", u.ID).Updates(&u)

	if result.Error != nil {
		return user.User{}, result.Error
	}

	if result.RowsAffected == 0 {
		return user.User{}, errors.New("record not found")
	}

	return u, nil
}

func (r UserRepository) Delete(id uint) error {
	result := initializers.DB.Delete(&user.User{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("record not found")
	}

	return nil
}
