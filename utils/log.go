package utils

// import (
// 	"log"

// 	"go.uber.org/zap"
// )

// var zapLogger zap.Logger

// func InitZap() {
// 	zapLogger, err := zap.NewProduction()
// 	if err != nil {
// 		log.Fatalf("can't initialize zap logger: %v", err)
// 	}
// 	defer zapLogger.Sync()
// }

// func Log(feature string, msg string) {
// 	zapLogger.Info(msg, zap.String("url", feature))
// }

// package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() {
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, os.Stdout, zapcore.InfoLevel)
	logg := zap.New(core, zap.AddCaller())

	zap.ReplaceGlobals(logg)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.LevelKey = "level"
	encoderConfig.TimeKey = "time"
	encoderConfig.MessageKey = "key"
	encoderConfig.EncodeTime = zapcore.TimeEncoder(func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		loc, err := time.LoadLocation("Asia/Bangkok")
		if err != nil {
			panic(err)
		}
		encoder.AppendString(t.In(loc).Format("02-Jan-2006 15:04:05"))
	})
	encoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewConsoleEncoder(encoderConfig)
}

// func LogInfo(msg string){
// 	zap.S().Info(msg)
// }

// func LogError (feature string msg string) {
// 	zap.S().Info(feature + msg)

// }
