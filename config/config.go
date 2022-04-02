package config

import (
	"context"
	"os"
)

type Config struct {
	Context   context.Context
	Port      string
	ProjectID string
	TopicID   string
	SubID     string
}

func GetConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	ctx := context.Background()
	return &Config{
		Context:   ctx,
		Port:      port,
		ProjectID: "kauche-practice",
		TopicID:   "test1",
		SubID:     "test1-sub",
	}
}
