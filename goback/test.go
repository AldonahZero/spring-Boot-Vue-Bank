package main
import (
    "github.com/gin-gonic/gin"
    "net/http"
)
func main() {
	// 默认服务器
    router:=gin.Default()
    router.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK,"Hello World")
    })
    router.Run(":8081")
}