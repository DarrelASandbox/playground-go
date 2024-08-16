package main_test

import (
	"context"
	"testing"

	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/gracefulshutdown/assert"
	specs_greet "github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/specs-greet"

	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/specs-greet/specifications"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestGreeterServer(t *testing.T) {
	ctx := context.Background()

	// Testcontainers gives us a programmatic way to build Docker images and manage container life-cycles.
	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:       "../../.",
			Dockerfile:    "./cmd/httpserver/Dockerfile",
			PrintBuildLog: true,
		},

		ExposedPorts: []string{"8080: 8080"},
		WaitingFor:   wait.ForHTTP("/").WithPort("8080"),
	}

	containers, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	assert.NoError(t, err)
	t.Cleanup(func() {
		assert.NoError(t, containers.Terminate(ctx))
	})

	driver := specs_greet.Driver{BaseURL: "http://localhost:8080"}
	specifications.GreetSpecification(t, driver)
}
