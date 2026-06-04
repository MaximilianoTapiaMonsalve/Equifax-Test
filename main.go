package main

import (
	"golang-interview/api/controllers"
	"golang-interview/domain/services"
	"golang-interview/infrastructure/https"
	"golang-interview/infrastructure/repositories"

	"github.com/gin-gonic/gin"
)

func main() {
	dummyClient := https.New("https://dummyjson.com")
	repo := repositories.New(dummyClient)
	service := services.New(repo)
	controller := controllers.New(service)

	r := gin.Default()
	r.GET("/dashboard/:id", controller.GetDashboard)
	r.Run(":3000") //port 8080 already in use
}
