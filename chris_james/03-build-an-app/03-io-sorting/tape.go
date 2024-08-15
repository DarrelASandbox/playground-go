package main

import "io"

/*
When we RecordWin, we Seek back to the start of the file and then write the new data—but
what if the new data was smaller than what was there before?
*/
type tape struct {
	file io.ReadWriteSeeker
}

func (t *tape) Write(p []byte) (n int, err error) {
	t.file.Seek(0, io.SeekStart)
	return t.file.Write(p)
}
