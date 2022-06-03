package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"title`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Learn Go", Completed: false},
	{ID: "2", Item: "Play Football", Completed: false},
	{ID: "3", Item: "Watch Anime", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func main() {
	server := gin.Default()

	//endpoint
	server.GET("/todos", getTodos)

	server.Run("localhost:8080")
}
