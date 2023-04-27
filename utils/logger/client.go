package logger

import "github.com/TskFok/AdminApi/global"

func Debug(debug interface{}) {
	defer global.LoggerClient.Sync()
	global.LoggerClient.Debug(debug)
}

func Error(error interface{}) {
	defer global.LoggerClient.Sync()
	global.LoggerClient.Error(error)
}

func Info(info interface{}) {
	defer global.LoggerClient.Sync()
	global.LoggerClient.Info(info)
}
