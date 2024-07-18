package main_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/specs-greet/adapters"
	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/specs-greet/adapters/httpserver"
	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/specs-greet/specifications"
)

func TestGreeterServer(t *testing.T) {
	var (
		host     = "localhost"
		port     = "8080"
		protocol = "tcp"
	)

	mappedPort := adapters.StartDockerServer(t, host, port, protocol, "httpserver")
	baseURL := fmt.Sprintf("http://%s:%s", host, mappedPort)
	driver := httpserver.Driver{BaseURL: baseURL, Client: &http.Client{Timeout: 1 * time.Second}}
	specifications.GreetSpecification(t, driver)
}
