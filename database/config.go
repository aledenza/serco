package database

import "time"

type DatabaseConfig struct {
	URL     string          `json:"url"`
	Options DatabaseOptions `json:"options"`
}

type DatabaseOptions struct {
	MaxIdleTime  time.Duration `json:"max_idle_time"  default:"0"`
	MaxLifeTime  time.Duration `json:"max_life_time"  default:"0"`
	MaxIdleConns int           `json:"max_idle_conns" default:"2"`
	MaxOpenConns int           `json:"max_open_conns" default:"0"`
}
