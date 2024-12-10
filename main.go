package main

import (
	"log"
	"time"
)

const (
	ConfigFilePath string        = "config.json"
	Timeout        time.Duration = 10000
)

func main() {
	ap, err := NewAutoPoster(ConfigFilePath, Timeout)
	if err != nil {
		log.Fatal(err)
	}
	ap.Run()
}
