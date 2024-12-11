package main

import (
	"fmt"
	"net/url"
	"time"
)

func Some[T any](arr []T, f func(T) bool) bool {
	for _, t := range arr {
		if f(t) {
			return true
		}
	}
	return false
}

func stripQueryString(u string) (string, error) {
	parsedUrl, err := url.Parse(u)
	if err != nil {
		return "", err
	}

	parsedUrl.RawQuery = ""

	return parsedUrl.String(), nil
}

func PrintWithTimestamp(s string) {
	timestamp := time.Now().Format(time.UnixDate)
	fmt.Printf("[%s] %s\n\n", timestamp, s)
}

func PrintWithTimestampf(format string, a ...any) {
	PrintWithTimestamp(fmt.Sprintf(format, a...))
}

func PrintErr(err error) {
	PrintWithTimestampf("Error: %s", err.Error())
}

func PrintErrWithTimeout(err error, timeout time.Duration) {
	PrintErr(err)
	Sleep(timeout)
}

func Sleep(d time.Duration) {
	PrintWithTimestampf("Sleeping for %f seconds", d.Seconds())
	time.Sleep(d)
}
