package webserver

import "github.com/go-rod/rod"

type Driver struct {
	baseURL string
	browser *rod.Browser
}

func NewDriver(baseURL string) (*Driver, func() error) {
	browser := rod.New().MustConnect()
	return &Driver{baseURL: baseURL, browser: browser}, browser.Close
}
