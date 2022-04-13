package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

func main() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	r.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})

	r.LoadHTMLGlob("templates/**/*")

	health := r.Group("/health")
	{
		health.GET("/ping", ping)
		health.GET("/status", status)
	}

	r.Run(":8080")
}

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
