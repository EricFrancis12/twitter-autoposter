package main

import (
	"fmt"
	"log"
	"time"
)

const (
	ConfigFilePath string        = "config.json"
	Timeout        time.Duration = time.Second * 10
)

func main() {
	ap, err := NewAutoPoster(ConfigFilePath, Timeout)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Autoposter Starting")
	ap.Run()
}
