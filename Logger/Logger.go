package Logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Logger *zap.Logger

func InitLogger() {
	file, _ := os.Create("./Logger/log/logger.log")
	writerSyncer := zapcore.AddSync(file)
	encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)

	Logger = zap.New(core)
}
