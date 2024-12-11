package main

import (
	"fmt"
	"net/url"
	"time"
)

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
	fmt.Printf("[%s] error: %s\n", timestamp, s)
}

func PrintWithTimestampf(format string, a ...any) {
	PrintWithTimestamp(fmt.Sprintf(format, a...))
}

func PrintErr(err error) {
	PrintWithTimestamp(err.Error())
}

func PrintErrWithTimeout(err error, timeout time.Duration) {
	PrintErr(err)
	Sleep(timeout)
}

func Sleep(d time.Duration) {
	PrintWithTimestampf("Sleeping for %f seconds", d.Seconds())
	time.Sleep(d)
}
