package main

import (
	"log"

	"github.com/RikiyaFujii/supporterz/src/database"
	"github.com/RikiyaFujii/supporterz/src/route"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// DBへの接続
	db, err := database.DBConnect()
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	defer db.Close()
	route.Routing(r, db)
	// ポート8000番をlisten
	r.Run(":8000")
}
