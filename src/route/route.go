package route

import (
	"github.com/RikiyaFujii/supporterz/src/controller"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Routing(r *gin.Engine, db *gorm.DB) {
	usersController := controller.NewUsersController(db)

	r.GET("/users", func(c *gin.Context) {
		usersController.LoadUsers(c)
	})

	r.GET("/users/:id", func(c *gin.Context) {
		usersController.LoadUser(c)
	})

	r.POST("/users", func(c *gin.Context) {
		usersController.AddUser(c)
	})

	r.PUT("/users/:id", func(c *gin.Context) {
		usersController.UpdateUser(c)
	})

	r.DELETE("/users/:id", func(c *gin.Context) {
		usersController.DeleteUser(c)
	})
}
