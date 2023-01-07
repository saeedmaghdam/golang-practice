package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := getUrlContentType("https://pkg.go.dev/static/shared/logo/go-white.svg")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Println(resp)
	}
}

func getUrlContentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	contentType := resp.Header.Get("content-type")
	if contentType == "" {
		return "", fmt.Errorf("can't find Content-Type header")
	}

	return contentType, err
}
