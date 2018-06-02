package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Todo struct
type Todo struct {
	ID   int
	Name string
}

// TodoIndex json
func TodoIndex(c *gin.Context) {
	todo := []Todo{
		{ID: 1, Name: "勉強"},
		{ID: 2, Name: "家事"},
		{ID: 3, Name: "愛ちゃん"},
	}
	c.JSON(http.StatusOK, todo)
}
