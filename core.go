package log

import "go.uber.org/zap/zapcore"

// Re-export zapcore types for backwards compatibility
type (
	EncoderConfig         = zapcore.EncoderConfig
	Encoder               = zapcore.Encoder
	StringDurationEncoder = zapcore.DurationEncoder
)

// Re-export zapcore functions
var (
	NewConsoleEncoder = zapcore.NewConsoleEncoder
)
