package server

import "github.com/gin-gonic/gin"

type ServerOptionFunc func(*Server)

func OptServerPort(port string) ServerOptionFunc {
	return func(s *Server) {
		s.port = port
	}
}

func OptServerMiddleware(middlewares ...gin.HandlerFunc) ServerOptionFunc {
	return func(s *Server) {
		s.middlewares = middlewares
	}
}
