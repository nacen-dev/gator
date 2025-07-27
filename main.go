package main

import (
	"fmt"
	"log"

	"github.com/nacen-dev/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	err = cfg.SetUser("vincent")
	if err != nil {
		log.Fatalf("couldn't set the user provided: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading the config: %v", err)
	}
	fmt.Printf("Read the config again: %+v\n", cfg)
}
