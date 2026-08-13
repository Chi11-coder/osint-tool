package main

import (
	"fmt"
	"os"
)

var VirusTotalKey string

func LoadVirusTotalKey() error {
	key, exists := os.LookupEnv("VIRUS_TOTAL")
	if !exists || key == "" {
		return fmt.Errorf("api key 'VIRUS_TOTAL' is not set")
	}

	VirusTotalKey = key
	return nil
}
