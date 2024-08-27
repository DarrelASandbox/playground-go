package poker_test

import (
	"io"
	"testing"

	poker "github.com/DarrelASandbox/playground-go/chris_james/03-build-an-app/websockets"
)

func TestTape_Write(t *testing.T) {
	file, clean := createTempFile(t, "12345")
	defer clean()
	tape := &poker.Tape{file}
	tape.Write([]byte("abc"))
	file.Seek(0, io.SeekStart)
	newFileContents, _ := io.ReadAll(file)
	got := string(newFileContents)
	want := "abc"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
