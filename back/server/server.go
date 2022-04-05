package server

import (
	"backend/db"

	"github.com/gin-gonic/gin"
)

type Server struct {
	r           *gin.Engine
	controller  *db.Controller
	middlewares []gin.HandlerFunc
	port        string
}

func NewServer(opts ...ServerOptionFunc) (*Server, error) {
	srv := &Server{}
	controller, err := db.NewController()
	if err != nil {
		return nil, err
	}
	srv.controller = controller
	for _, optFunc := range opts {
		optFunc(srv)
	}
	srv.r = srv.router()

	return srv, nil
}

func (srv *Server) router() *gin.Engine {
	r := gin.Default()
	r.Use(srv.middlewares...)
	r.POST("/json/todos", srv.handlePostTodo)
	r.GET("/json/todos", srv.handleGetAllTodos)
	r.GET("/json/todos/:id", srv.handleGetTodoByID)
	r.GET("/text/todos", srv.handleGetAllTodosText)
	r.GET("/text/todos/:id", srv.handleGetTodoByIDText)
	return r
}

func (srv *Server) Run() error {
	return srv.r.Run(":" + srv.port)
}
