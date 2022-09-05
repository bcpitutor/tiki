package tlogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func InitiliazeTikiLogger() {
	loggerOutputFile := "./tiki.log"

	conf := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		OutputPaths: []string{loggerOutputFile, "stdout"},

		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "msg",
			LevelKey:     "level",
			TimeKey:      "time",
			CallerKey:    "file",
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	Log, _ = conf.Build()
}
