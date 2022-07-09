# 一个文件实现一个静态文件服务器

> 使用gin

```go
package main

import (
	"fmt"
	"net/http"
	"os"

	"flag"

	"github.com/gin-gonic/gin"
)

var (
	port *string
	dir  *string
)

func main() {

	port = flag.String("port", "9393", "http port")
	dir = flag.String("dir", "./out/", "static file dir")
	flag.Parse()

	router := gin.Default()
	staticWorkDir := *dir

	router.StaticFile("/", staticWorkDir+"/home/stat.html")
	router.StaticFS("/_next", http.Dir(staticWorkDir+"/_next"))
	router.NoRoute(renderPage)
	router.Run(":" + *port)

}

func renderPage(ctx *gin.Context) {
	renderHtml(ctx, ctx.Request.URL.Path)
}

func renderHtml(ctx *gin.Context, fileName string) {
	ctx.Header("Cache-Control", "private, max-age=86400")

	fullPath := fmt.Sprintf("%s%s.html", *dir, fileName)
	_, err := os.Stat(fullPath)
	if err == nil {
		content, _ := os.ReadFile(fullPath)
		ctx.Data(http.StatusOK, "text/html", content)
		return
	}
	ctx.Redirect(http.StatusFound, "/")
}

```