package interfaces

import "github.com/NetKBs/uneg-sase-go-backend/api/models/user"

type IRoleRepository interface {
	Create(role user.Role) (user.Role, error)
	Get(id uint) (user.Role, error)
	GetAll() ([]user.Role, error)
	Update(role user.Role) (user.Role, error)
	Delete(id uint) error
}
