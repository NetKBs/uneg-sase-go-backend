package sql

import (
	"github.com/NetKBs/uneg-sase-go-backend/api/models/user"
	"github.com/NetKBs/uneg-sase-go-backend/initializers"
)

type RoleRepositoryImpl struct {
}

func (r RoleRepositoryImpl) Create(role user.Role) (user.Role, error) {

	result := initializers.DB.Create(&role)

	if result.Error != nil {
		return user.Role{}, result.Error
	}

	return user.Role{Id: role.Id, Name: role.Name}, nil
}
