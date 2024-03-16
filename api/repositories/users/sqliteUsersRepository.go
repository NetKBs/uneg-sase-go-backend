package users

import (
	"github.com/NetKBs/uneg-sase-go-backend/api/models"
	"gorm.io/gorm"
)

type SQLiteUsersRepository struct {
	DB *gorm.DB
}

func NewSQLiteUsersRepository(db *gorm.DB) *SQLiteUsersRepository {
	return &SQLiteUsersRepository{DB: db}
}

func (r *SQLiteUsersRepository) FindAll() ([]models.User, error) {
	return nil, nil
}

func (r *SQLiteUsersRepository) FindByID(id int) (models.User, error) {
	return models.User{}, nil
}

func (r *SQLiteUsersRepository) Create(user models.User) (models.User, error) {
	return user, nil
}

func (r *SQLiteUsersRepository) Update(user models.User) (models.User, error) {
	return user, nil
}

func (r *SQLiteUsersRepository) Delete(id int) error {
	return nil
}
