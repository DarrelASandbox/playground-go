package cancelreader

import (
	"io"
)

/*
--- FAIL: TestContextAwareReader (0.00s)
    --- FAIL: TestContextAwareReader/behaves_like_a_normal_reader (0.00s)
panic: runtime error: invalid memory address or nil pointer dereference [recovered]
        panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x2 addr=0x0 pc=0x102a0fd78]
*/
func NewCancellableReader(rdr io.Reader) io.Reader {
	return nil
}
