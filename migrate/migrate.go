package main

import (
	"sheik.com/go/crud/initializers"
	"sheik.com/go/crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&model.Post{})
}