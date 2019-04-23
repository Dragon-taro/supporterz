package controller

import (
	"net/http"

	"github.com/RikiyaFujii/supporterz/src/common"
	"github.com/RikiyaFujii/supporterz/src/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type UsersController struct {
	DB *gorm.DB
}

func NewUsersController(db *gorm.DB) *UsersController {
	u := new(UsersController)
	u.DB = db
	return u
}

func (u *UsersController) LoadUsers(c *gin.Context) {
	users, err := model.LoadUsers(u.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, users)
}

func (u *UsersController) LoadUser(c *gin.Context) {
	id := common.ToInt(c)
	user, err := model.LoadUser(u.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, user)
}

func (u *UsersController) AddUser(c *gin.Context) {
	name, email := common.PostFormBuilder(c)

	if err := model.AddUser(u.DB, name, email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

func (u *UsersController) UpdateUser(c *gin.Context) {
	name, email := common.PostFormBuilder(c)
	id := common.ToInt(c)

	if err := model.UpdateUser(u.DB, id, name, email); err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

func (u *UsersController) DeleteUser(c *gin.Context) {
	id := common.ToInt(c)

	if err := model.DeleteUser(u.DB, id); err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
