package api

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	apiGroup := r.Group("/api")
	apiGroup.GET("/ping", ping)
}
