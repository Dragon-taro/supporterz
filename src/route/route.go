package route

import (
	"github.com/RikiyaFujii/supporterz/src/controller"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Routing(r *gin.Engine, db *gorm.DB) {
	// インスタンス
	usersController := controller.NewUsersController(db)
	// apiのパスのグルーピング
	v1 := r.Group("/api/v1")
	{
		v1.GET("/users", func(c *gin.Context) {
			usersController.LoadUsers(c)
		})

		v1.GET("/users/:id", func(c *gin.Context) {
			usersController.LoadUser(c)
		})

		v1.POST("/users", func(c *gin.Context) {
			usersController.AddUser(c)
		})

		v1.PUT("/users/:id", func(c *gin.Context) {
			usersController.UpdateUser(c)
		})

		v1.DELETE("/users/:id", func(c *gin.Context) {
			usersController.DeleteUser(c)
		})
	}

}
