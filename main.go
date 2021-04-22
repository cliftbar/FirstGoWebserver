package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"site.cliftbar/FirstWebserver/controllers"
)

func main(){
	fmt.Println("hello")

	r := gin.Default()

	// Define controllers
	health := new(controllers.HealthController)
	ping := new(controllers.PingController)

	// Define routes
	r.GET("/health", health.Status)
	r.GET("/ping", ping.Ping)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}