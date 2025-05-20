package app

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()

	// Register routes (akan kita buat nanti)
	// RegisterRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(fmt.Sprintf(":%s", port))
}
