package main

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	NatsAddress string `envconfig:"NATS_ADDRESS"`
	NatsToken   string `envconfig:"NATS_TOKEN"`
	Port        int    `envconfig:"PORT"`
	StreamName  string `envconfig:"STREAM_NAME"`
}

func main() {
	cfg := &Config{}
	if err := envconfig.Process("", cfg); err != nil {
		log.Fatalf("Could not parse env Config: %v", err)
	}

	fmt.Printf("%+v", cfg)
}
