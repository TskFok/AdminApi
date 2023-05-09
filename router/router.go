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

	Handle.GET("/corpus", controller.CorpusVicunaList)
	Handle.POST("/corpus", controller.AddCorpusVicuna)
	Handle.PUT("/corpus", controller.UpdateCorpusVicuna)
	Handle.DELETE("/corpus/:id", controller.DelCorpusVicuna)
}
