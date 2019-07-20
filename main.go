package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	g := gin.New()

	http.ListenAndServe("localhost:8080",g)
}
