package sql

import (
	"errors"

	"github.com/NetKBs/uneg-sase-go-backend/api/models/user"
	"github.com/NetKBs/uneg-sase-go-backend/initializers"
)

type RoleRepository struct {
}

func (r RoleRepository) Create(role user.Role) (user.Role, error) {

	result := initializers.DB.Create(&role)

	if result.Error != nil {
		return user.Role{}, result.Error
	}

	return user.Role{Id: role.Id, Name: role.Name}, nil
}

func (r RoleRepository) Get(id uint) (user.Role, error) {
	var role user.Role
	result := initializers.DB.First(&role, id)

	if result.Error != nil {
		return user.Role{}, result.Error
	}

	return user.Role{Id: role.Id, Name: role.Name}, nil
}

func (r RoleRepository) GetAll() ([]user.Role, error) {
	var roles []user.Role
	result := initializers.DB.Find(&roles)

	if result.Error != nil {
		return []user.Role{}, result.Error
	}

	return roles, nil
}

func (r RoleRepository) Update(role user.Role) (user.Role, error) {

	result := initializers.DB.Model(&role).Where("Id = ?", role.Id).Updates(&role)

	if result.Error != nil {
		return user.Role{}, result.Error
	}

	if result.RowsAffected == 0 {
		return user.Role{}, errors.New("record not found")
	}

	return user.Role{Id: role.Id, Name: role.Name}, nil

}

func (r RoleRepository) Delete(id uint) error {
	result := initializers.DB.Delete(&user.Role{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("record not found")
	}

	return nil
}
