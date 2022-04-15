package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func status(c *gin.Context) {

	c.HTML(http.StatusOK, "health/status.tmpl", gin.H{
		"status": "OK",
		"now":    time.Date(2022, 01, 01, 0, 0, 0, 0, time.UTC),
	})
}

func dbStatus(c *gin.Context) {
	healthChecks, _ := GetHealthChecks(10)

	c.HTML(http.StatusOK, "health/db.tmpl", gin.H{
		"data": healthChecks,
	})
}
