package user

import (
	"net/http"
	"strconv"

	"github.com/NetKBs/uneg-sase-go-backend/api/models/user"
	repository "github.com/NetKBs/uneg-sase-go-backend/api/repositories/user/sql"
	service "github.com/NetKBs/uneg-sase-go-backend/api/services/user"
	"github.com/gin-gonic/gin"
)

// CreateRole godoc
// @Summary Creates a new user role.
// @Description Create a new user role.
// @Accept json
// @Produce json
// @Param role body user.RoleInputCreate true "RoleInputCreate"
// @Success 200 {object} user.Role
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /users/roles/create [post]
func Create(c *gin.Context) {
	var dto user.RoleInputCreate

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	repo := &repository.RoleRepository{}
	service := service.NewRoleService(repo)

	result, err := service.Create(user.Role{
		Name: dto.Name,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// UpdateRole godoc

// UpdateRole godoc
// @Summary Updates a user role.
// @Description Updates a user role.
// @Accept json
// @Produce json
// @Param role body user.RoleInputUpdate true "RoleInputUpdate"
// @Success 200 {object} user.Role
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /users/roles/update [put]

func Update(c *gin.Context) {

	var dto user.RoleInputUpdate

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	repo := &repository.RoleRepository{}
	service := service.NewRoleService(repo)

	result, err := service.Update(user.Role(dto))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// Delete godoc
// @Summary Delete a role by ID
// @Description Delete a role by its ID
// @Accept json
// @Produce json
// @Param id path int true "Role ID"
// @Success 200 {object} gin.H{"message": "deleted"}
// @Failure 400 {object} gin.H{"error": "id should be a number"}
// @Failure 500 {object} gin.H{"error": "Internal Server Error"}
// @Router /users/roles/update [delete]

func Delete(c *gin.Context) {

	id := c.Param("id")

	if _, err := strconv.Atoi(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id should be a number"})
		return
	}

	id_num, _ := strconv.Atoi(id)

	repo := &repository.RoleRepository{}
	service := service.NewRoleService(repo)

	if err := service.Delete(uint(id_num)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})

}

// GetAll godoc
// @Summary Get all roles
// @Description Get all roles
// @Accept json
// @Produce json
// @Success 200 {object} YourResultType
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/roles [get]

func GetAll(c *gin.Context) {
	repo := &repository.RoleRepository{}
	service := service.NewRoleService(repo)

	result, err := service.GetAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// Get godoc
// @Summary Get a role by ID
// @Description Get a role by its ID
// @Accept json
// @Produce json
// @Param id path int true "Role ID"
// @Success 200 {object} YourResultType
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/roles/{id} [get]

func Get(c *gin.Context) {

	id := c.Param("id")

	if _, err := strconv.Atoi(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id should be a number"})
		return
	}

	id_num, _ := strconv.Atoi(id)

	repo := &repository.RoleRepository{}
	service := service.NewRoleService(repo)

	result, err := service.Get(uint(id_num))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)

}
