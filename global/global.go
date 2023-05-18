package global

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var AppMode string
var RedisUser string
var RedisPassword string
var RedisHost string
var MysqlDsn string
var MysqlPrefix string
var JwtSecret string
var JwtExpire int
var RedisClient *redis.Client
var DataBase *gorm.DB
var LoggerFilePath string
var LoggerClient *zap.SugaredLogger
var OpenAiToken string
var VicunaUrl string
