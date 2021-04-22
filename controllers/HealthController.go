package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"site.cliftbar/FirstWebserver/services"
	"time"
)

type HealthController struct{}
type PingController struct{}

func (h HealthController) Status(c *gin.Context) {
	healthCheck := services.HealthCheck()
	c.String(http.StatusOK, healthCheck)
}

func (h PingController) Ping(c *gin.Context) {
	pingResponse := services.Ping()
	time.LoadLocation("US/Pacific")
	c.String(http.StatusOK, pingResponse.Format(time.RFC3339))
}