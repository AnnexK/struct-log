package main

import (
	"errors"

	"github.com/AnnexK/struct-log/internal/config"
	zapimpl "github.com/AnnexK/struct-log/internal/infrastructure/logging/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogObject struct {
	NestedField1 string
	NestedField2 int
	Ok bool
}

func (obj LogObject) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("nestedField1", obj.NestedField1)
	enc.AddInt("nestedField2", obj.NestedField2)
	enc.AddBool("ok", obj.Ok)
	return nil
}

type ReflectedLogObject struct {
	Field1 string `json:"field1"`
	Field2 int `json:"field2"`
	Ok bool `json:"ok"`
}

func main() {
	cfg := config.MustLoadConfig()
	_ = zapimpl.InitLogging(cfg)
	defer zap.L().Sync()
	
	// Basic logging
	zap.L().Info("hello zap", zap.Int("someField", 64), zap.String("someOtherField", "kek"))

	// Errors
	zap.L().Error("error for zap", zap.Error(errors.New("something bad happened")))

	// Logger context
	logger := zap.L().With(zap.String("loggerField", "loggerField"), zap.Int("otherLoggerField", 321))
	logger.Info("hello zap", zap.Int("someField", 64))

	// Nested objects
	zap.L().Info("hello zap", zap.Object("logObject", LogObject{
		NestedField1: "neField",
		NestedField2: 166,
		Ok: true,
	}))

	// Nested objects (Reflection)
	zap.L().Info("hello zap", zap.Reflect("logObject", ReflectedLogObject{
		Field1: "loggerField",
		Field2: 164,
		Ok: true,
	}))
}
