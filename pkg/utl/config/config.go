package config

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Load returns Configuration struct
func Load(path string) (*Configuration, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}
	var cfg = new(Configuration)
	if err := yaml.Unmarshal(bytes, cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}
	return cfg, nil
}

// Configuration holds data necessery for configuring application
type Configuration struct {
	Server *Server      `yaml:"server,omitempty"`
	DB     *Database    `yaml:"database,omitempty"`
	App    *Application `yaml:"application,omitempty"`
}

// Database holds data necessery for database configuration
type Database struct {
	PSN     string `yaml:"psn,omitempty"`
	Timeout int    `yaml:"timeout_seconds,omitempty"`
}

// Server holds data necessery for server configuration
type Server struct {
	Port         string `yaml:"port,omitempty"`
	Debug        bool   `yaml:"debug,omitempty"`
	ReadTimeout  int    `yaml:"read_timeout_seconds,omitempty"`
	WriteTimeout int    `yaml:"write_timeout_seconds,omitempty"`
}

// Application holds application configuration details
type Application struct {
	SwaggerUIPath string `yaml:"swagger_ui_path,omitempty"`
}
