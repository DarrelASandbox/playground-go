package main

import (
	"testing"
	"time"

	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/gracefulshutdown/acceptancetests"
	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/gracefulshutdown/assert"
)

const (
	port = "8081"
	url  = "http://localhost:" + port
)

func TestNonGracefulShutdown(t *testing.T) {
	cleanup, sendInterrupt, err := acceptancetests.LaunchTestProgram(port)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(cleanup)

	// just check the server works before we shut things down
	assert.CanGet(t, url)
	time.AfterFunc(50*time.Millisecond, func() {
		assert.NoError(t, sendInterrupt())
	})
	assert.CantGet(t, url)
}
