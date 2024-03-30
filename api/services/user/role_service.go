package user

import (
	"github.com/NetKBs/uneg-sase-go-backend/api/models/user"
	"github.com/NetKBs/uneg-sase-go-backend/api/repositories/user/interfaces"
)

type RoleService struct {
	repository interfaces.RoleRepository
}

func NewRoleService(repository interfaces.RoleRepository) *RoleService {
	return &RoleService{repository}
}

func (r *RoleService) Create(role user.Role) (user.Role, error) {
	return r.repository.Create(role)
}
