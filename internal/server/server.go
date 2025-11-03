package server

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	return gin.Default()
}

func NewRawRouter() *gin.Engine {
	return gin.New()
}
