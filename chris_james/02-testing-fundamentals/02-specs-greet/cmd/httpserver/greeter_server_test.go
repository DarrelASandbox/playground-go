package main_test

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/specs-greet/adapters/httpserver"
	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/specs-greet/specifications"
	"github.com/alecthomas/assert/v2"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

/*
When you expose a port in a Docker container (e.g., 8080/tcp), Docker maps this port to a port on the host system.
If you don’t explicitly specify which host port to use, Docker will choose a random, available port on the host machine.
*/
func TestGreeterServer(t *testing.T) {
	ctx := context.Background()

	// Testcontainers gives us a programmatic way to build Docker images and manage container life-cycles.
	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:       "../../.",
			Dockerfile:    "./cmd/httpserver/Dockerfile",
			PrintBuildLog: true,
		},

		ExposedPorts: []string{"8080/tcp"},
		WaitingFor:   wait.ForListeningPort("8080/tcp"),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	mappedPort, errPort := container.MappedPort(ctx, "8080/tcp")
	if errPort != nil {
		t.Fatalf("Failed to get mapped port: %v", errPort)
	}

	log.Printf("Container is running on port %s", mappedPort.Port())

	logs, errTC := container.Logs(ctx)
	if errTC != nil {
		log.Fatalf("Failed to get logs: %v", errTC)
	}
	io.Copy(os.Stdout, logs)

	assert.NoError(t, err)
	t.Cleanup(func() {
		assert.NoError(t, container.Terminate(ctx))
	})

	client := http.Client{Timeout: 1 * time.Second}

	driver := httpserver.Driver{
		BaseURL: fmt.Sprintf("http://localhost:%s", mappedPort.Port()),
		Client:  &client,
	}
	specifications.GreetSpecification(t, driver)
}
