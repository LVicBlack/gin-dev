package sync

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	apiGroup := r.Group("/sync")
	apiGroup.GET("/ping", ping)
}
