package cancelreader

import (
	"context"
	"io"
)

func NewCancellableReader(ctx context.Context, rdr io.Reader) io.Reader {
	return rdr
}
