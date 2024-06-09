package logger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zapSugarLog *zap.SugaredLogger
var trackingFieldKeys []string
var ZapBasicLogger *zap.Logger

type LogConfig struct {
	LogLevel             string
	ServiceName          string
	TrackingFieldKeysReq []string
}

func init() {
	Init(&LogConfig{})
}

// Init overrides the base init if required for OS Environment & Exclusion.
func Init(lc *LogConfig) {
	trackingFieldKeys = lc.TrackingFieldKeysReq
	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.StacktraceKey = ""
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig = encoderConfig

	var globalLoggingLevel zapcore.Level

	switch lc.LogLevel {
	case "DEBUG":
		globalLoggingLevel = zapcore.DebugLevel
	case "INFO":
		globalLoggingLevel = zapcore.InfoLevel
	case "WARN":
		globalLoggingLevel = zapcore.WarnLevel
	case "ERROR":
		globalLoggingLevel = zapcore.ErrorLevel
	default:
		globalLoggingLevel = zapcore.DebugLevel
	}

	config.Level = zap.NewAtomicLevelAt(globalLoggingLevel)

	if lc.ServiceName != "" {
		config.InitialFields = map[string]interface{}{
			"engine": lc.ServiceName,
		}
	} else {
		config.InitialFields = map[string]interface{}{
			"engine": "not_specified",
		}
	}

	// Caller Skip allows the proper caller instead of showing pkg caller
	ZapLogger, err := config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	// clears buffers if any
	defer ZapLogger.Sync()
	zapSugarLog = ZapLogger.Sugar()
}

// InitBasic initializes a sugared logger instance with basic configuration.
func InitBasic(lc *LogConfig) {
	var err error
	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.StacktraceKey = ""
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig = encoderConfig

	var globalLoggingLevel zapcore.Level

	switch lc.LogLevel {
	case "DEBUG":
		globalLoggingLevel = zapcore.DebugLevel
	case "INFO":
		globalLoggingLevel = zapcore.InfoLevel
	case "WARN":
		globalLoggingLevel = zapcore.WarnLevel
	case "ERROR":
		globalLoggingLevel = zapcore.ErrorLevel
	default:
		globalLoggingLevel = zapcore.DebugLevel
	}

	config.Level = zap.NewAtomicLevelAt(globalLoggingLevel)

	if lc.ServiceName != "" {
		config.InitialFields = map[string]interface{}{
			"engine": lc.ServiceName,
		}
	} else {
		config.InitialFields = map[string]interface{}{
			"engine": "not_specified",
		}
	}

	// Caller Skip allows the proper caller instead of showing pkg caller
	ZapBasicLogger, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	// clears buffers if any
	defer ZapBasicLogger.Sync()
	zapSugarLog = ZapBasicLogger.Sugar()
}

// Debug logs a message at DEBUG level with some additional context from fields.
func Debug(msg string, ctx context.Context, fields map[string]interface{}) {
	var keyValueField []interface{}
	for k, v := range fields {
		keyValueField = append(keyValueField, k, v)
	}

	keyValueField = getTrackingDetails(ctx, keyValueField)

	zapSugarLog.Debugw(msg, keyValueField...)
}

// Info logs a message at INFO level with some additional context from fields
func Info(msg string, ctx context.Context, fields map[string]interface{}) {
	var keyValueField []interface{}
	for k, v := range fields {
		keyValueField = append(keyValueField, k, v)
	}

	keyValueField = getTrackingDetails(ctx, keyValueField)

	zapSugarLog.Infow(msg, keyValueField...)
}

// Warn logs a message at WARN level with some additional context from fields.
func Warn(msg string, ctx context.Context, fields map[string]interface{}) {
	var keyValueField []interface{}
	for k, v := range fields {
		keyValueField = append(keyValueField, k, v)
	}

	keyValueField = getTrackingDetails(ctx, keyValueField)

	zapSugarLog.Warnw(msg, keyValueField...)
}

// Error logs a message at ERROR level with some additional context from fields.
func Error(msg string, ctx context.Context, fields map[string]interface{}) {
	var keyValueField []interface{}
	for k, v := range fields {
		keyValueField = append(keyValueField, k, v)
	}

	keyValueField = getTrackingDetails(ctx, keyValueField)

	zapSugarLog.Errorw(msg, keyValueField...)
}

// getTrackingDetails checks whether ctx has values defined by keys in trackingFieldKeys.
// If yes, it appends the key value to keyValueField to be logged.
func getTrackingDetails(ctx context.Context, keyValueField []interface{}) []interface{} {
	for _, key := range trackingFieldKeys {
		if val, exist := ctx.Value(key).(interface{}); exist {
			keyValueField = append(keyValueField, key, val)
		}
	}

	return keyValueField
}