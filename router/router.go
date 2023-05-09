package router

import (
	"github.com/TskFok/AdminApi/controller"
	"github.com/TskFok/AdminApi/global"
	"github.com/TskFok/AdminApi/middleware"
	"github.com/gin-gonic/gin"
)

var Handle *gin.Engine

func InitRouter() {
	gin.SetMode(global.AppMode)

	Handle = gin.New()
	Handle.Use(gin.Recovery())
	Handle.Use(gin.Logger())
	Handle.Use(middleware.Cors())

	Handle.POST("/login", controller.Login)

	Handle.Use(middleware.Jwt())

	Handle.GET("/home", controller.HomeData)
	Handle.GET("/user", controller.UserList)
	Handle.POST("/user", controller.AddUser)
	Handle.PUT("/user/status", controller.UpdateStatus)
	Handle.GET("/corpus", controller.CorpusList)
	Handle.POST("/corpus", controller.AddCorpus)
	Handle.PUT("/corpus", controller.UpdateCorpus)
	Handle.DELETE("/corpus/:id", controller.DelCorpus)

	Handle.GET("/corpus-vicuna", controller.CorpusVicunaList)
	Handle.POST("/corpus-vicuna", controller.AddCorpusVicuna)
	Handle.PUT("/corpus-vicuna", controller.UpdateCorpusVicuna)
	Handle.DELETE("/corpus-vicuna/:id", controller.DelCorpusVicuna)
}
