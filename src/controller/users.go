package controller

import (
	"net/http"

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
