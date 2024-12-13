package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

const (
	ConfigFilePath           string  = "config.json"
	DefaultMinTimeoutMinutes float64 = 30
	DefaultMaxTimeoutMinutes float64 = 60
	DefaultErrTimeoutMinutes float64 = 0.2
)

func main() {
	var (
		minTimeoutFlag = flag.Float64("mint", DefaultMinTimeoutMinutes, "Minimum timeout in minutes")
		maxTimeoutFlag = flag.Float64("maxt", DefaultMaxTimeoutMinutes, "Maximum timeout in minutes")
		errTimeoutFlag = flag.Float64("errt", DefaultErrTimeoutMinutes, "Error timeout in minutes")
	)
	flag.Parse()

	var (
		minTimeout = time.Duration(*minTimeoutFlag * float64(time.Minute))
		maxTimeout = time.Duration(*maxTimeoutFlag * float64(time.Minute))
		errTimeout = time.Duration(*errTimeoutFlag * float64(time.Minute))
	)

	if minTimeout > maxTimeout {
		log.Fatal(fmt.Errorf("minimum timeout must be less than maximum timeout"))
	}

	ap, err := NewAutoPoster(ConfigFilePath, minTimeout, maxTimeout, errTimeout)
	if err != nil {
		log.Fatal(err)
	}

	PrintWithTimestamp("Autoposter Starting")
	ap.Run()
}
