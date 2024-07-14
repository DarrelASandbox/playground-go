package adapters

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/gracefulshutdown/assert"
	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func StartDockerServer(
	t testing.TB,
	host string,
	port string,
	protocol string,
	dockerFilePath string,
) string {
	exposedPorts := fmt.Sprintf("%s/%s", port, protocol)
	ctx := context.Background()
	t.Helper()

	// Testcontainers gives us a programmatic way to build Docker images and manage container life-cycles.
	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:       "../../.",
			Dockerfile:    dockerFilePath,
			PrintBuildLog: true,
		},

		ExposedPorts: []string{exposedPorts},
		WaitingFor:   wait.ForListeningPort(nat.Port(exposedPorts)),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	mappedPort, errPort := container.MappedPort(ctx, nat.Port(exposedPorts))
	if errPort != nil {
		t.Fatalf("Failed to get mapped port: %v", errPort)
	}

	log.Printf("Container is running on port %s", mappedPort.Port())

	assert.NoError(t, err)
	t.Cleanup(func() {
		assert.NoError(t, container.Terminate(ctx))
	})

	return mappedPort.Port()
}
