package main

import (
	"os"
	"time"

	clockface "github.com/DarrelASandbox/playground-go/chris_james/16-clockface"
)

/*
# Build steps
cd clockface
go build
./clockface > clock.svg
cd ..
*/

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
