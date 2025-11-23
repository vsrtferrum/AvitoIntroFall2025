package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/vsrtferrum/AvitoIntroFall2025/internal/infrastructure/storage"
)

func ReadConfig(logger Logger) (*storage.Config, error) {
	var config config
	err := cleanenv.ReadEnv(&config)
	if err != nil {
		logger.WriteError(err)
		return nil, ErrFailedToReadConfig
	}
	return &storage.Config{
		Host:              config.Host,
		Port:              config.Port,
		User:              config.User,
		Password:          config.Password,
		Name:              config.Name,
		PoolSize:          config.PoolSize,
		HealthCheckPeriod: config.HealthCheckPeriod,
		MaxConnLifetime:   config.MaxConnLifetime,
		MaxConnIdleTime:   config.MaxConnIdleTime,
	}, nil
}
