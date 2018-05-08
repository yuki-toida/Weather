package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("ui/templates/**/*")

	router.GET("/", func(context *gin.Context) {
		html := template.Must(template.ParseFiles("ui/templates/layout/main.html", "ui/templates/home/index.html"))
		router.SetHTMLTemplate(html)
		context.HTML(http.StatusOK, "main.html", nil)
	})
	router.Run()
}
