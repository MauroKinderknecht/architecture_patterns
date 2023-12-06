package logger

import (
	"github.com/MauroKinderknecht/tech_talk/microservices/item_service/config"
	"go.uber.org/zap"
)

type Logger struct {
	*zap.Logger
}

func (l *Logger) With(fields ...zap.Field) *Logger {
	return &Logger{
		Logger: l.Logger.With(fields...),
	}
}

func String(key string, val string) zap.Field {
	return zap.String(key, val)
}

func Error(err error) zap.Field {
	return zap.Error(err)
}

func NewLogger(config *config.Config) (*Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = config.LogCfg.Output

	if config.LogCfg.Level != "" {
		level, err := zap.ParseAtomicLevel(config.LogCfg.Level)
		if err != nil {
			return nil, err
		}
		cfg.Level = level
	}

	log, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return &Logger{
		Logger: log,
	}, nil
}
