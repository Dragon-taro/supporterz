package common

import "github.com/gin-gonic/gin"

func PostFormBuilder(c *gin.Context) (string, string) {
	return c.PostForm("name"), c.PostForm("email")
}
