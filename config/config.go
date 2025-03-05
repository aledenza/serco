package config

import (
	"github.com/aledenza/serco/client"
	"github.com/aledenza/serco/database"
)

type Config struct {
	DbConf     database.DatabaseConfig `json:"db_conf"`
	ClientConf client.ClientConfig     `json:"client_conf"`
	Token      string                  `json:"token"`
}
