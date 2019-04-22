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

	router.GET("/users/:id", func(c *gin.Context) {
		usersController.LoadUser(c)
	})

	router.POST("/users", func(c *gin.Context) {
		usersController.AddUser(c)
	})

	router.PUT("/users/:id", func(c *gin.Context) {
		usersController.UpdateUser(c)
	})

	router.DELETE("/users/:id", func(c *gin.Context) {
		usersController.DeleteUser(c)
	})
}
