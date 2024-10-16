package config

import (
	"github.com/BurntSushi/toml"
)

type Host struct {
	UpdateUrl string `toml:"update_url"`
	Username  string `toml:"username"`
	Password  string `toml:"password"`
	Hostname  string `toml:"hostname"`
	IP        string
	Offline   string `toml:"offline"`
}

type Config struct {
	Hosts []Host
}

func GetConfig(config_file string) (Config, error) {
	// parse TOML and return config
	var config Config
	_, err := toml.DecodeFile(config_file, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
