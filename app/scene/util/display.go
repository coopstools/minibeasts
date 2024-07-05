package util

import (
	"fmt"
	"io"
)

type GameCtx struct {
	io.Writer
	io.Reader
}

func ListOptions(ctx GameCtx, opts []string) {
	for i, name := range opts {
		_, _ = fmt.Fprintf(ctx, "%d: %s\n", i, name)
	}
}
