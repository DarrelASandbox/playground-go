package main_test

import (
	"context"
	"testing"
	"time"

	specs_greet "github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/specs-greet"

	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/specs-greet/specifications"
	"github.com/alecthomas/assert/v2"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

/*
All requested ports were not exposed: port 8080:8080 is not mapped yet
2024/08/16 12:50:21 ⏳ Waiting for container id fe0d7a51d794 image: 6699ef05-45a0-4ae8-93b1-20d6239fffea:6027fe5d-ddbb-43e1-a53f-0b9b2b79e955.
Waiting for: &{timeout:0x1400046ccc0 Port:8080
Path:/ StatusCodeMatcher:0x1043724e0
ResponseMatcher:0x104456b50
UseTLS:false AllowInsecure:false
TLSConfig:<nil> Method:GET Body:<nil>
Headers:map[] ResponseHeadersMatcher:0x104456b60
PollInterval:100ms UserInfo: ForceIPv4LocalHost:false}

2024/08/16 12:50:21 container logs (all exposed ports, [8080:8080], were not mapped in 5s: port 8080:8080 is not mapped yet):

	greeter_server_test.go:36: Did not expect an error but got:
	    failed to start container: all exposed ports, [8080:8080], were not mapped in 5s: port 8080:8080 is not mapped yet

--- FAIL: TestGreeterServer (29.53s)
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

		ExposedPorts: []string{"8080:8080"},
		WaitingFor:   wait.ForHTTP("/").WithPort("8080").WithStartupTimeout(30 * time.Second),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	assert.NoError(t, err)
	t.Cleanup(func() {
		assert.NoError(t, container.Terminate(ctx))
	})

	driver := specs_greet.Driver{BaseURL: "http://localhost:8080"}
	specifications.GreetSpecification(t, driver)
}
