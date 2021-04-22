package services

import "time"

func HealthCheck() (status string) {
	status = "healthy"
	return
}

func Ping() (pingResponse time.Time) {
	pingResponse = time.Now().UTC()
	return
}
