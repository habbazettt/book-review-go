package main

import (
	"github.com/habbazettt/book-review-go/internal/app"
	"github.com/habbazettt/book-review-go/internal/config"
)

func main() {
	config.LoadEnv()
	config.InitLogger()
	config.ConnectDB()
	defer config.CloseDB()

	app.Run()
}
