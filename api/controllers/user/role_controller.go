package user

import (
	"net/http"

	"github.com/NetKBs/uneg-sase-go-backend/api/models/user"
	repository "github.com/NetKBs/uneg-sase-go-backend/api/repositories/user/sql"
	service "github.com/NetKBs/uneg-sase-go-backend/api/services/user"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var dto user.RoleInputCreate

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	repo := &repository.RoleRepositoryImpl{}
	service := service.NewRoleService(repo)

	result, err := service.Create(user.Role{Name: dto.Name})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)

}
