package main

import (
	"fmt"

	"github.com/crackeer/goaway/container"
	_ "github.com/crackeer/goaway/examples/sign" // 自定义实现签名算法
	"github.com/crackeer/goaway/server"
	"github.com/gin-gonic/gin"
)

func init() {
	// 自定义注册路由
	container.RegisterNakedRouter("access/token", func(ctx *gin.Context) {
		fmt.Println("access/token")
	})
}
func main() {
	server.Main()
}
