package main

import (
	"log"

	"github.com/RikiyaFujii/supporterz/src/database"
	"github.com/RikiyaFujii/supporterz/src/route"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db, err := database.DBConnect()
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	defer db.Close()

	route.Routing(router, db)
	router.Run(":8080")
}
