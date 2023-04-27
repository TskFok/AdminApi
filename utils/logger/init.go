package logger

import (
	"bytes"
	"fmt"
	"github.com/TskFok/AdminApi/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func InitLogger() *zap.SugaredLogger {
	path := bytes.NewBufferString(global.LoggerFilePath)

	dirErr := os.Mkdir(global.LoggerFilePath, 0755)

	if nil != dirErr {
		fmt.Println(dirErr.Error())
	}
	path.WriteString(global.AppMode)
	path.WriteString(".log")

	logPath := path.String()

	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)

	if err != nil {
		fmt.Println(err)
	}

	if nil != err {
		fmt.Println(err.Error())
	}
	sync := zapcore.AddSync(file)
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	core := zapcore.NewCore(encoder, sync, zap.DebugLevel)
	logger := zap.New(core)
	sugarLogger := logger.Sugar()

	return sugarLogger
}
