package main

import (
	"fmt"
	"time"
)

func PrintErr(err error) {
	timestamp := time.Now().Format(time.UnixDate)
	fmt.Printf("[%s] error: %s\n", timestamp, err.Error())
}

func PrintErrWithTimeout(err error, timeout time.Duration) {
	PrintErr(err)
	time.Sleep(timeout)
}
