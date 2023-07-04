package config

import "github.com/spf13/viper"

// MustLoadConfig загружает конфиг.
func MustLoadConfig() (config Config) {
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	if err := tryReadConfig(&config); err != nil {
		panic(err)
	}

	return
}

func tryReadConfig(config *Config) error {
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(config); err != nil {
		return err
	}
	return nil
}
