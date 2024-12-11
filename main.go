package main

import (
	"log"
	"time"
)

const (
	ConfigFilePath string = "config.json"
	// TODO: Impliment a max and min timeout to reduce chance of bot detection,
	// and also make it configurable via flags.
	Timeout time.Duration = time.Minute * 18
)

func main() {
	ap, err := NewAutoPoster(ConfigFilePath, Timeout)
	if err != nil {
		log.Fatal(err)
	}
	PrintWithTimestamp("Autoposter Starting")
	ap.Run()
}
