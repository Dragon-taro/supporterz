package main

import (
	"log"

	"github.com/RikiyaFujii/supporterz/src/database"
	"github.com/RikiyaFujii/supporterz/src/route"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db, err := database.DBConnect()
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	defer db.Close()
	route.Routing(r, db)
	r.Run(":8000")
}
