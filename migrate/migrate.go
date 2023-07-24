package main

import (
	"github.com/gaurangkudale/go-crud/initializers"
	"github.com/gaurangkudale/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
