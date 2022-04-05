package server

import (
	"backend/db"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (src *Server) handlePostTodo(ctx *gin.Context) {
	todo := &db.Todo{}
	if err := ctx.BindJSON(todo); err != nil {
		log.Println(err)
		return
	}
	newTodo, err := src.controller.CreateNewTodo(todo)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, newTodo)
}

func (src *Server) handleGetAllTodos(ctx *gin.Context) {
	todos, err := src.controller.FetchAllTodos()
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, todos)
}

func (src *Server) handleGetTodoByID(ctx *gin.Context) {
	idString, ok := ctx.Params.Get("id")
	if !ok {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	todo, err := src.controller.FetchTodoByID(id)
	if err != nil {
		log.Println(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatus(http.StatusNotFound)
		} else {
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
		return
	}
	ctx.JSON(http.StatusOK, todo)
}
