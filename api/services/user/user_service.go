package user

import (
	"github.com/NetKBs/uneg-sase-go-backend/api/models/user"
	"github.com/NetKBs/uneg-sase-go-backend/api/repositories/user/interfaces"
)

type UserService struct {
	repository interfaces.IUserRepository
}

func NewUserService(respository interfaces.IUserRepository) *UserService {
	return &UserService{respository}
}

func (r *UserService) Create(user user.User) (user.User, error) {
	return r.repository.Create(user)
}

func (r *UserService) Update(user user.User) (user.User, error) {
	return r.repository.Update(user)
}

func (r *UserService) Get(id uint) (user.User, error) {
	return r.repository.Get(id)
}

func (r *UserService) GetAll() ([]user.User, error) {
	return r.repository.GetAll()
}

func (r *UserService) Delete(id uint) error {
	return r.repository.Delete(id)
}
