package user

import (
	"net/http"

	"github.com/NetKBs/uneg-sase-go-backend/api/models"
	"github.com/NetKBs/uneg-sase-go-backend/api/models/user"
	"github.com/gin-gonic/gin"

	repository "github.com/NetKBs/uneg-sase-go-backend/api/repositories/user/sql"
	service "github.com/NetKBs/uneg-sase-go-backend/api/services/user"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
	var body user.UserInputCreate

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Error: err.Error()})
		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Error: err.Error()})
		return
	}

	body.Password = string(hash)

	repo := &repository.UserRepository{}
	service := service.NewUserService(repo)

	result, err := service.Create(user.User{Username: body.Username, Password: body.Password, RoleId: body.RoleId})

	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetAllUser(c *gin.Context) {
	repo := &repository.UserRepository{}
	service := service.NewUserService(repo)

	result, err := service.GetAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
