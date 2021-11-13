package server

import (
	"github.com/gin-gonic/gin"
)

// init gin as http server
func NewHTTPServer(greeter *service.) (*Engine, error) {
	router := gin.Default()
	router.Run(":8080")
	return &router, nil
}
