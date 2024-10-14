package routers

import "github.com/gin-gonic/gin"

type Option func(engine *gin.Engine)

var options = []Option{}

func Register(option ...Option) {
	options = append(options, option...)
}

func Init(router *gin.Engine) {
	for _, opt := range options {
		opt(router)
	}
}
