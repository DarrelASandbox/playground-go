package osexec

import (
	"encoding/xml"
	"os/exec"
	"strings"
	"testing"
)

type Payload struct {
	Message string `xml:"message"`
}

// The problem with GetData is the business logic is coupled with the means of getting the XML.
// To make our design better we need to decouple them
func GetData() string {
	cmd := exec.Command("cat", "msg.xml")
	out, _ := cmd.StdoutPipe()
	var payload Payload
	decoder := xml.NewDecoder(out)

	// these 3 can return errors but I'm ignoring for brevity
	cmd.Start()
	decoder.Decode(&payload)
	cmd.Wait()
	return strings.ToUpper(payload.Message)
}

func TestGetData(t *testing.T) {
	got := GetData()
	want := "HAPPY NEW YEAR!"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
