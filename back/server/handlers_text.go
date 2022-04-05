package server

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (src *Server) handleGetAllTodosText(ctx *gin.Context) {
	todos, err := src.controller.FetchAllTodos()
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	jsonText, err := json.Marshal(todos)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.Data(http.StatusOK, "plain/text", jsonText)
}

func (src *Server) handleGetTodoByIDText(ctx *gin.Context) {
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
	jsonText, err := json.Marshal(todo)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.Data(http.StatusOK, "plain/text", jsonText)
}
