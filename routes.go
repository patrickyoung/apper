package main

import "github.com/gin-gonic/gin"

func routes(r *gin.Engine) {
	health := r.Group("/health")
	{
		health.GET("/ping", ping)
		health.GET("/status", status)
		health.GET("/db", dbStatus)
	}
}
