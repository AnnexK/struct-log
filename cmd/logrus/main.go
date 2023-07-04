package main

import (
	"errors"

	"github.com/AnnexK/struct-log/internal/config"
	logrusimpl "github.com/AnnexK/struct-log/internal/infrastructure/logging/logrus"
	"github.com/sirupsen/logrus"
)

type FieldHook struct {
	fields map[string]any
}

func NewFieldHook(fields map[string]any) *FieldHook {
	return &FieldHook{
		fields: fields}
}

func (hook *FieldHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
		logrus.TraceLevel,
	}
}

func (hook *FieldHook) Fire(entry *logrus.Entry) error {
	for k, v := range hook.fields {
		entry.WithField(k, v)
	}
	return nil
}

type NestedObject struct {
	NestedField1 string `json:"nestedField1"`
	NestedField2 int    `json:"nestedField2"`
	Ok           bool   `json:"ok"`
}

func main() {
	cfg := config.MustLoadConfig()
	_ = logrusimpl.InitLogging(cfg)

	// Basic logging
	// Note: only reflection
	logrus.StandardLogger().WithField("someField", 64).WithField("someOtherField", "kek").Info("hello logrus")

	// Errors
	logrus.StandardLogger().WithError(errors.New("something bad happened")).Error("error for logrus")

	// No out-of-the-box logger context

	// Nested objects -- reflection only
	logrus.StandardLogger().WithField("nObj", NestedObject{
		NestedField1: "neField",
		NestedField2: 165,
		Ok: true,
	}).Info("hello logrus")
}
