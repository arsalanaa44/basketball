package main

import (
	"fmt"
	"log"

	"github.com/arsalanaa44/basketball/initializers"
	"github.com/arsalanaa44/basketball/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.User{}, &models.Basket{})
	fmt.Println("? Migration complete")
}
