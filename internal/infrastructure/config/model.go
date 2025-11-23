package config

import "time"

type config struct {
	Host              string        `env:"PG_HOST" envDefault:"localhost"`
	User              string        `env:"PG_USER" envDefault:"admin"`
	Password          string        `env:"PG_PASSWORD" envDefault:"admin"`// да это небезопасно но у vault кубера тоже нет 
	Name              string        `env:"PG_DB" envDefault:"avitointro"`
	Port              int           `env:"PG_PORT" envDefault:"5432"`
	HealthCheckPeriod time.Duration `env:"PG_HEALTH_CHECK_PERIOD" envDefault:"30s"`
	MaxConnLifetime   time.Duration `env:"PG_MAX_CONN_LIFETIME" envDefault:"1h"`
	MaxConnIdleTime   time.Duration `env:"PG_MAX_CONN_IDLE_TIME" envDefault:"5m"`
	PoolSize          int32         `env:"PG_POOL_SIZE" envDefault:"10"`
}
