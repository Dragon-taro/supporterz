package route

import (
	"github.com/RikiyaFujii/supporterz/src/controller"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Routing(router *gin.Engine, db *gorm.DB) {
	usersController := controller.NewUsersController(db)

	router.GET("/users", func(c *gin.Context) {
		usersController.LoadUsers(c)
	})
}
