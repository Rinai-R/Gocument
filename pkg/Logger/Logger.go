package Logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Logger *zap.Logger

func init() {
	file, _ := os.OpenFile("./pkg/Logger/log/logger.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	writerSyncer := zapcore.AddSync(file)
	encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)

	Logger = zap.New(core)
}
