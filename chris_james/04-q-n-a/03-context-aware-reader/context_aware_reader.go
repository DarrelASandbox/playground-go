package cancelreader

import (
	"context"
	"io"
)

type readerCtx struct {
	ctx      context.Context
	delegate io.Reader
}

// NewCancellableReader will stop reading to rdr if ctx is cancelled.
func NewCancellableReader(ctx context.Context, rdr io.Reader) io.Reader {
	return &readerCtx{ctx: ctx, delegate: rdr}
}

func (r *readerCtx) Read(p []byte) (n int, err error) {
	if err := r.ctx.Err(); err != nil {
		return 0, err
	}

	return r.delegate.Read(p)
}
