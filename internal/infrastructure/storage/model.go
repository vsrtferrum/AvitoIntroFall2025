package storage

import "time"

type Config struct {
	Host              string
	User              string
	Password          string
	Name              string
	Port              int
	HealthCheckPeriod time.Duration
	MaxConnLifetime   time.Duration
	MaxConnIdleTime   time.Duration
	PoolSize          int32
}
