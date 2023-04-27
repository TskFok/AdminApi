package bootstrap

import (
	"github.com/TskFok/AdminApi/global"
	"github.com/TskFok/AdminApi/utils/cache"
	"github.com/TskFok/AdminApi/utils/conf"
	"github.com/TskFok/AdminApi/utils/database"
	"github.com/TskFok/AdminApi/utils/logger"
	"os"
)

func Init() {
	//守护进程
	args := os.Args

	if len(args) != 1 && args[1] == "bg" {
		initProcess()
	}

	//配置文件
	conf.InitConfig()

	//redis
	global.RedisClient = cache.InitRedis()
	//mysql
	global.DataBase = database.InitMysql()
	//log
	global.LoggerClient = logger.InitLogger()
}
