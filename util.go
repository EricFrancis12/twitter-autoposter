package main

import "net/url"

func stripQueryString(u string) (string, error) {
	parsedUrl, err := url.Parse(u)
	if err != nil {
		return "", err
	}

	parsedUrl.RawQuery = ""

	return parsedUrl.String(), nil
}
