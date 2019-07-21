package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	//g.Run("localhost:8080")

	// gin middlewares
	//middlewares := []gin.HandlerFunc{}

	//

	log.Printf("start to listening the incoming request on port %s", "8080")
	log.Printf(http.ListenAndServe("localhost:8080", g).Error())
}
