package users

import "github.com/NetKBs/uneg-sase-go-backend/api/models"

type UsersRepository interface {
	FindAll() ([]models.User, error)
	FindByID(id int) (models.User, error)
	Create(user models.User) (models.User, error)
	Update(user models.User) (models.User, error)
	Delete(id int) error
}
