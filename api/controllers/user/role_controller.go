package user

import (
	"net/http"
	"strconv"

	"github.com/NetKBs/uneg-sase-go-backend/api/models"
	"github.com/NetKBs/uneg-sase-go-backend/api/models/user"
	repository "github.com/NetKBs/uneg-sase-go-backend/api/repositories/user/sql"
	service "github.com/NetKBs/uneg-sase-go-backend/api/services/user"
	"github.com/gin-gonic/gin"
)

// CreateRole godoc
//
//	@summary		create new role
//	@description	create a new role with a name
//	@tags			roles
//	@accept			json
//	@produce		json
//	@param role body user.RoleInputCreate true "Role Input"
//	@success		200	{object} user.Role
//	@failure		400	{object} models.Error
//	@router			/users/roles/create [post]
func Create(c *gin.Context) {
	var dto user.RoleInputCreate

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Error: err.Error()})
		return
	}

	repo := &repository.RoleRepository{}
	service := service.NewRoleService(repo)

	result, err := service.Create(user.Role{
		Name: dto.Name,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// UpdateRole godoc
//
//	@summary		update an existing role
//	@description	update an existing role with a new name
//	@tags			roles
//	@accept			json
//	@produce		json
//	@param 			role body user.Role true "Role Input"
//	@success		200	{object} user.Role
//	@failure		400	{object} models.Error
//	@router			/users/roles/update [put]
func Update(c *gin.Context) {

	var dto user.RoleInputUpdate

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Error: err.Error()})
		return
	}

	repo := &repository.RoleRepository{}
	service := service.NewRoleService(repo)

	result, err := service.Update(user.Role(dto))

	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// DeleteRole godoc
//
//	@summary		delete a role
//	@description	delete a role by id
//	@tags			roles
//	@produce		json
//	@param 			id path int true "Role ID"
//	@success		200	{object} models.Messages
//	@failure		400	{object} models.Error
//	@router			/users/roles/delete/{id} [delete]
func Delete(c *gin.Context) {

	id := c.Param("id")

	if _, err := strconv.Atoi(id); err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Error: "id should be a number"})
		return
	}

	id_num, _ := strconv.Atoi(id)

	repo := &repository.RoleRepository{}
	service := service.NewRoleService(repo)

	if err := service.Delete(uint(id_num)); err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.Messages{Message: "deleted"})

}

// GetAllRoles godoc
//
//	@summary		get all roles
//	@description	return the data of all roles existing
//	@tags			roles
//	@produce		json
//	@success		200	{object} user.Role
//	@failure		400	{object} models.Error
//	@router			/users/roles/getall [get]
func GetAll(c *gin.Context) {
	repo := &repository.RoleRepository{}
	service := service.NewRoleService(repo)

	result, err := service.GetAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetRole godoc
//
//	@summary		get a role
//	@description	get a role by id
//	@tags			roles
//	@produce		json
//	@param 			id path int true "Role ID"
//	@success		200	{object} user.Role
//	@failure		400	{object} models.Error
//	@router			/users/roles/get/{id} [get]
func Get(c *gin.Context) {

	id := c.Param("id")

	if _, err := strconv.Atoi(id); err != nil {
		c.JSON(http.StatusBadRequest, models.Error{Error: "id should be a number"})
		return
	}

	id_num, _ := strconv.Atoi(id)

	repo := &repository.RoleRepository{}
	service := service.NewRoleService(repo)

	result, err := service.Get(uint(id_num))

	if err != nil {
		c.JSON(http.StatusNotFound, models.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)

}
