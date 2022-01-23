package config

import "go.uber.org/dig"

type Config struct {
}

func NewConfig() *Config {
	return &Config{}
}

func Init(container *dig.Container) {
	container.Provide(NewConfig)
}
