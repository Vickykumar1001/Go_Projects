package handlers

import (
	"net/http"
	"user_management/internal/dto"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllUsers(c *gin.Context) {

	users, err := h.repo.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed To Get users",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    users,
	})

}
func (h *Handler) GetUser(c *gin.Context) {
	username := c.Param("username")

	user, err := h.repo.GetUserByUserName(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	username := c.Param("username")

	var updatedUsr dto.SignupRequest
	err := c.BindJSON(&updatedUsr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Request",
		})
		return
	}
	user, err := h.repo.GetUserByUserName(username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	(*user).Password = updatedUsr.Password
	(*user).Name = updatedUsr.Name
	(*user).Email = updatedUsr.Email

	h.repo.UpdateUser(username, *user)

	c.JSON(http.StatusOK, gin.H{
		"message": "User Updated",
		"user":    user,
	})
}

func (h *Handler) DeleteUser(c *gin.Context) {
	username := c.Param("username")

	err := h.repo.DeleteUser(username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
