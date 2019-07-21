package conf

import (
	"github.com/BurntSushi/toml"
)

var Conf *Cfg

type (
	Cfg struct {
		Log        *LogCfg        `toml:"Log"`
		PostgreSQL *PostgreSQLCfg `toml:"PostgreSQL"`
	}

	LogCfg struct {
		LogLevel string `toml:"log_level"`
		FilePath string `toml:"file_path"`
	}

	PostgreSQLCfg struct {
		Host     string `toml:"host"`
		Port     string `toml:"port"`
		Username string `toml:"username"`
		Password string `toml:"password"`
	}
)

func Init() error {
	if _, err := toml.DecodeFile("conf/conf.toml", &Conf); err != nil {
		panic(err)
	}
	return nil
}
