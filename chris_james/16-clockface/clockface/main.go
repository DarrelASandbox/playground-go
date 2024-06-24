package main

import (
	"os"
	"time"

	svg "github.com/DarrelASandbox/playground-go/chris_james/16-clockface/svg"
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
	svg.SVGWriter(os.Stdout, t)
}
