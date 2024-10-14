package main

import (
	"gin-dev/app/api"
	"gin-dev/app/sync"
	"gin-dev/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routers.Register(api.Router, sync.Router)
	routers.Init(r)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
