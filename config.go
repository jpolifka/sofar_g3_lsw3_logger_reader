package main

import (
	"errors"
	"github.com/kubaceg/sofar_g3_lsw3_logger_reader/adapters/export/otlp"
	"os"

	"github.com/kubaceg/sofar_g3_lsw3_logger_reader/adapters/export/mosquitto"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Inverter struct {
		Port         string `yaml:"port"`
		LoggerSerial uint   `yaml:"loggerSerial"`
		ReadInterval int    `default:"60" yaml:"readInterval"`
	} `yaml:"inverter"`
	Mqtt mosquitto.MqttConfig `yaml:"mqtt"`
	Otlp otlp.Config          `yaml:"otlp"`
}

func (c *Config) validate() error {
	if c.Inverter.Port == "" {
		return errors.New("missing required inverter.port config")
	}

	if c.Inverter.LoggerSerial == 0 {
		return errors.New("missing required inverter.loggerSerial config")
	}

	return nil
}

func NewConfig(configPath string) (*Config, error) {
	config := &Config{}
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	if err := config.validate(); err != nil {
		return nil, err
	}

	return config, nil
}
