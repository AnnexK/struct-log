package main

import (
	"errors"

	"github.com/AnnexK/struct-log/internal/config"
	zlogimpl "github.com/AnnexK/struct-log/internal/infrastructure/logging/zerolog"
	zlog "github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type NestedObject struct {
	NestedField1 string `json:"nestedField1"`
	NestedField2 int    `json:"nestedField2"`
	Ok           bool   `json:"ok"`
}

func main() {
	cfg := config.MustLoadConfig()
	_ = zlogimpl.InitLogging(cfg)

	// Basic logging
	log.Info().Int("someField", 64).Str("someOtherField", "kek").Msg("hello zerolog")

	// Errors
	log.Error().Err(errors.New("something bad happened")).Msg("error for zerolog")

	// Logger context
	logger := log.With().Str("loggerField", "loggerField").Int("otherLoggerField", 321).Logger()
	logger.Info().Int("someField", 64).Msg("hello zerolog")

	// Nested objects (Dict)
	log.Info().Dict("dict", zlog.Dict().Str("nestedField1", "neField").Int("nestedField2", 165).
		Bool("ok", true)).Msg("hello zerolog")

	// Nested objects (Reflection)
	log.Info().Any("obj", NestedObject{
		NestedField1: "neField",
		NestedField2: 165,
		Ok:           true,
	}).Msg("hello zerolog")
}
