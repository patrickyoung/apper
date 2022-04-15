package main

import (
	"embed"
	"fmt"
	"html/template"
	"time"

	"github.com/gin-gonic/gin"
)

//go:embed templates/*
var f embed.FS

func templateFunctions(r *gin.Engine) {
	myFuncMap := template.FuncMap{
		"formatAsDate": formatAsDate,
	}

	// Load templates from embedded filesystem
	templ := template.Must(template.New("").Funcs(myFuncMap).ParseFS(f, "templates/**/*"))
	r.SetHTMLTemplate(templ)
}

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}
