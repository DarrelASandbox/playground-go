package main_test

import (
	"fmt"
	"testing"

	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/specs-greet/adapters"
	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/specs-greet/adapters/webserver"
	"github.com/alecthomas/assert/v2"
)

func TestGreeterWeb(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	var (
		host     = "localhost"
		port     = "8081"
		protocol = "tcp"
	)

	mappedPort := adapters.StartDockerServer(t, host, port, protocol, "webserver")
	baseURL := fmt.Sprintf("http://%s:%s", host, mappedPort)
	t.Logf("mappedPort: %s", mappedPort)
	_, cleanup := webserver.NewDriver(baseURL)
	webserver.NewDriver(baseURL)
	t.Cleanup(func() { assert.NoError(t, cleanup()) })
}
