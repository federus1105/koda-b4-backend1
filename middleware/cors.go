package middleware

import (
	"net/http"
	"os"
	"slices"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware(ctx *gin.Context) {
	origin := ctx.GetHeader("Origin")
	
	whitelist := []string{
		os.Getenv("ORIGIN_URL"),
	}
	if slices.Contains(whitelist, origin) {
		ctx.Header("Access-Control-Allow-Origin", origin)
	}

	ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
	ctx.Header("Access-Control-Allow-Headers", "Authorization, Content-Type")

	if ctx.Request.Method == http.MethodOptions {
		ctx.AbortWithStatus(http.StatusNoContent)
		return
	}
	ctx.Next()
}
