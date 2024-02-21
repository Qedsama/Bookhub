package main

import (
    "fmt"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func main() {
    r := gin.Default()

    config := cors.DefaultConfig()
    config.AllowOrigins = []string{"*"}  // 允许所有来源，你也可以指定具体的域名
    config.AllowMethods = []string{"*"}
    r.Use(cors.New(config))


    r.POST("/api/register", registerHandler)
    r.POST("/api/login", loginHandler)
    // 其他路由

    fmt.Println("Server is running on :28000")
    r.Run(":28000")
}