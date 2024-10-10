package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("hello name: %s, age: %d ", "TipGo", 21)

	// //logger
	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "TipsGo"), zap.Int("age", 21))

	// logger := zap.NewExample()
	// logger.Info("Hello")

	// //Development
	// logger, _ = zap.NewDevelopment()
	// logger.Info(("Hello NewDevelopment"))

	// //Production
	// logger, _ = zap.NewProduction()
	// logger.Info("Hello NewProduction")

	//Customer
	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 2))
}

// format logs a message
func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	//convert 1728542076.9310627 -> 2024-10-10T13:34:36.930+0700
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//ts -> Time
	encodeConfig.TimeKey = "time"
	//from info -> INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	//"caller":"cli/main.log.go:26"
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encodeConfig)
}

func getWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
