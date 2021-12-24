package config

import "github.com/BurntSushi/toml"

type Config struct {
	Server Server `toml:"server"`
	Store  Store  `toml:"database"`
}

type Server struct {
	Addr       string `toml:"addr"`
	LogLevel   int    `toml:"loglevel"`
	SessionKey string `toml:"session_key"`
	Debug      bool   `toml:"debug"`
}

type Store struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	DBName   string `toml:"dbname"`
}

// New Config from toml file
func New(configFile string) (*Config, error) {
	config := &Config{}
	_, err := toml.DecodeFile(configFile, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
