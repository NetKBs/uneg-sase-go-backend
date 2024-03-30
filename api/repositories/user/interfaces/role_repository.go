package interfaces

import "github.com/NetKBs/uneg-sase-go-backend/api/models/user"

type RoleRepository interface {
	Create(role user.Role) (user.Role, error)
}
