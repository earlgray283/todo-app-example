package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello from gin-gonic!")
	})
	r.POST("/todos", func(ctx *gin.Context) {
		newTodo := &Todo{}
		if err := ctx.BindJSON(newTodo); err != nil {
			log.Println(err)
			return
		}
		newTodo.createdAt = time.Now()
		todoStorage.Lock()
		defer todoStorage.Unlock()
		todoStorage.Todos = append(todoStorage.Todos, *newTodo)
	})
	r.GET("/todos/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		todo, ok := todoMap[id]
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No such todo which id is %v", id),
			})
			return
		}
		ctx.JSON(http.StatusOK, &todo)
	})
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
