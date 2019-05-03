package common

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PostFormBuilder(c *gin.Context) (string, string) {
	// body からの値を返している
	return c.PostForm("name"), c.PostForm("email")
}

func ToInt(c *gin.Context) int {
	// stringからintに変換
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	return id
}
