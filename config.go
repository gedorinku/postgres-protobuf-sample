package main

import "github.com/creasty/configo"

type Config struct {
	DataBaseURL string `envconfig:"database_url" valid:"required"`
}

func loadConfig() (*Config, error) {
	c := &Config{}
	if err := configo.Load(c, configo.Option{}); err != nil {
		return nil, err
	}
	return c, nil
}
