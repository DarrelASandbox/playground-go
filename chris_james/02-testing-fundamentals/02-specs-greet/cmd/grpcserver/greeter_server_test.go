package main_test

import (
	"fmt"
	"testing"

	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/specs-greet/adapters"
	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/specs-greet/adapters/grpcserver"
	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/specs-greet/specifications"
)

func TestGreeterServer(t *testing.T) {
	var (
		host     = "localhost"
		port     = "50051"
		protocol = "tcp"
	)

	mappedPort := adapters.StartDockerServer(t, host, port, protocol, "grpcserver")
	driver := grpcserver.Driver{Addr: fmt.Sprintf("%s:%s", host, mappedPort)}
	specifications.GreetSpecification(t, &driver)
}
