package user

import (
	"github.com/NetKBs/uneg-sase-go-backend/api/models/user"
	"github.com/NetKBs/uneg-sase-go-backend/api/repositories/user/interfaces"
)

type RoleService struct {
	repository interfaces.IRoleRepository
}

func NewRoleService(repository interfaces.IRoleRepository) *RoleService {
	return &RoleService{repository}
}

func (r *RoleService) Create(role user.Role) (user.Role, error) {
	return r.repository.Create(role)
}

func (r *RoleService) Get(id uint) (user.Role, error) {
	return r.repository.Get(id)
}

func (r *RoleService) GetAll() ([]user.Role, error) {
	return r.repository.GetAll()
}

func (r *RoleService) Update(role user.Role) (user.Role, error) {
	return r.repository.Update(role)
}

func (r *RoleService) Delete(id uint) error {
	return r.repository.Delete(id)
}
