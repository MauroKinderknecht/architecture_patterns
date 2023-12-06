package config

import (
	"github.com/spf13/viper"
	"os"
	"strings"
)

// Config stores the configuration for the backend.
type Config struct {
	ApiPort  int          `mapstructure:"api_port"`
	MongoURL string       `mapstructure:"mongodb_url"`
	LogCfg   LoggerConfig `mapstructure:"log"`
}

type LoggerConfig struct {
	Level  string   `mapstructure:"level"`
	Output []string `mapstructure:"output"`
}

func LoadConfig() (*Config, error) {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "development"
	}

	viper.AddConfigPath("./config")
	viper.SetConfigName(env)
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()

	cfg := Config{}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
