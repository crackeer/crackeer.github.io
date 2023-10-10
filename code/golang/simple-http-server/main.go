package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := "8888"
	if len(os.Args) >= 2 {
		port = os.Args[1]
	}
	router := gin.Default()
	router.StaticFS("/", http.Dir("./"))

	router.Run(":" + port)
}
