package utils

import (
	"path/filepath"
	"os"
	"encoding/json"
	"errors"
)

type Config struct {
	Name, DisplayName, Description string

	Dir  string
	Exec string
	Args []string
	Env  []string

	Stderr, Stdout string
}

func GetConfigPath() (string, error) {
	if E != "" {
		return filepath.Join(E, C), nil
	}
	return "", errors.New("Config Path is unreachable")
}

func GetConfig(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	conf := &Config{}

	r := json.NewDecoder(f)
	err = r.Decode(&conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
