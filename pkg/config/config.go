package config

import (
	"errors"
	"fmt"

	"github.com/BurntSushi/toml"
)

type Host struct {
	UpdateUrl string `toml:"update_url"`
	Username  string `toml:"username"`
	Password  string `toml:"password"`
	Hostname  string `toml:"hostname"`
	IP        string `toml:"ip"`
	Offline   string `toml:"offline"`
}

type Config struct {
	Hosts []Host
}

func GetConfig(config_file string) (Config, error) {
	// parse TOML and return config
	var config Config
	tomlData, err := toml.DecodeFile(config_file, &config)
	if err != nil {
		return config, err
	}

	// check for missing configuration values
	for i, host := range config.Hosts {
		switch {
		case host.Hostname == "":
			return config, errors.New(fmt.Sprintf("Missing 'hostname' in host #%d in configuration file", i+1))
		case host.UpdateUrl == "":
			return config, errors.New(fmt.Sprintf("Missing 'update_url' in host #%d in configuration file", i+1))
		case host.Username == "":
			return config, errors.New(fmt.Sprintf("Missing 'username' in host #%d in configuration file", i+1))
		case host.Password == "":
			return config, errors.New(fmt.Sprintf("Missing 'password' in host #%d in configuration file", i+1))
		case host.Offline == "":
			return config, errors.New(fmt.Sprintf("Missing 'offline' in host #%d in configuration file", i+1))
		}
	}

	// check for extra configuration values
	undecoded := tomlData.Undecoded()
	if len(undecoded) != 0 {
		for _, u := range undecoded {
			fmt.Printf("Not recognized entry in configuration file: %s\n", u)
		}
		return config, errors.New("Extra data detected")
	}

	return config, nil
}
