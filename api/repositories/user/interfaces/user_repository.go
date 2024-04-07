package interfaces

import "github.com/NetKBs/uneg-sase-go-backend/api/models/user"

type IUserRepository interface {
	Create(user user.User) (user.User, error)
	Get(id uint) (user.User, error)
	GetAll() ([]user.User, error)
	Update(role user.User) (user.User, error)
	Delete(id uint) error
}
