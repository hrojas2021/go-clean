package logger

import (
	"os"
	"time"

	"github.com/hugo.rojas/custom-api/conf"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZapLogger(c *conf.LoggerConfiguration, conf zap.Config, isLocal bool) *zap.Logger {
	conf.EncoderConfig = setCustomEncoder()
	syncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   c.DefaultPath,
		MaxSize:    c.MaxSize,
		MaxBackups: c.MaxBackups,
		MaxAge:     c.MaxAge,
	})

	if isLocal {
		syncer = zap.CombineWriteSyncers(os.Stdout, syncer)
	}

	logRotate := zapcore.NewCore(
		zapcore.NewJSONEncoder(conf.EncoderConfig),
		syncer,
		conf.Level,
	)

	l := zap.New(zapcore.NewTee(logRotate), zap.AddCaller())
	return l
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02T15:04:05"))
}

func setCustomEncoder() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:  "Msg",
		LevelKey:    "Lvl",
		TimeKey:     "Time",
		NameKey:     "Layer",
		EncodeTime:  customTimeEncoder,
		EncodeLevel: zapcore.LowercaseLevelEncoder,
	}
}
