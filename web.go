package main

import (
	"embed"
	"html/template"

	"github.com/gin-gonic/gin"
)

//go:embed templates/*
var f embed.FS

func initWeb() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	templateFunctions(r)
	routes(r)

	r.Run(":8080")
}

func templateFunctions(r *gin.Engine) {
	myFuncMap := template.FuncMap{
		"formatAsDate": formatAsDate,
	}

	// Load templates from embedded filesystem
	templ := template.Must(template.New("").Funcs(myFuncMap).ParseFS(f, "templates/**/*"))
	r.SetHTMLTemplate(templ)
}

func routes(r *gin.Engine) {
	health := r.Group("/health")
	{
		health.GET("/ping", ping)
		health.GET("/status", status)
	}
}
