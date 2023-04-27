package conf

import (
	"bytes"
	_ "embed"
	"github.com/TskFok/AdminApi/global"
	"github.com/spf13/viper"
)

//go:embed conf.yaml
var conf []byte

func InitConfig() {
	viper.SetConfigType("yaml")

	err := viper.ReadConfig(bytes.NewReader(conf))

	if nil != err {
		panic(err)
	}

	global.AppMode = viper.Get("app.mode").(string)
	global.RedisUser = viper.Get("redis.user").(string)
	global.RedisPassword = viper.Get("redis.password").(string)
	global.RedisHost = viper.Get("redis.host").(string)
	global.MysqlDsn = viper.Get("mysql.dsn").(string)
	global.MysqlPrefix = viper.Get("mysql.prefix").(string)
	global.JwtSecret = viper.Get("jwt.secret").(string)
	global.JwtExpire = viper.Get("jwt.expire").(int)
	global.LoggerFilePath = viper.Get("logger.file_path").(string)
	global.OpenAiToken = viper.Get("openai.token").(string)
}
