package logger

import (
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(w io.Writer) *zap.Logger {
	priority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return true
	})

	wr := zapcore.AddSync(w)
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, wr, priority),
	)

	return zap.New(core)
}
