package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	g := gin.Default()

	//g.Run("localhost:8080")
	http.ListenAndServe("localhost:8080", g)
}
