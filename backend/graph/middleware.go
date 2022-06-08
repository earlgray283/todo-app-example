package graph

import (
	"context"
	"net/http"

	firebase "firebase.google.com/go/v4"
	"github.com/gin-gonic/gin"
)

func MiddlewareSessionCookie() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		newCtx := context.WithValue(ctx.Request.Context(), respWriterKey, ctx.Writer)
		ctx.Request = ctx.Request.WithContext(newCtx)
		ctx.Next()
	}
}

func MiddlewareAuth(fb *firebase.App) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session, err := ctx.Cookie("session")
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		client, err := fb.Auth(ctx)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		token, err := client.VerifySessionCookie(ctx, session)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		user, _ := client.GetUser(ctx, token.UID)
		newCtx := context.WithValue(ctx.Request.Context(), userKey, user)
		ctx.Request = ctx.Request.WithContext(newCtx)
		ctx.Next()
	}
}
