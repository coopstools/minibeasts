package util

import (
	"fmt"
	"io"
)

const (
	listItem = "%d: %s\n"
)

type GameCtx struct {
	io.Writer
	io.Reader
}

func (c GameCtx) Print(t string) {
	_, _ = fmt.Fprint(c, t)
}

func (c GameCtx) Printf(format string, args ...any) {
	_, _ = fmt.Fprintf(c, format, args...)
}

func ListOptions(ctx GameCtx, opts []string) {
	for i, name := range opts {
		ctx.Printf(listItem, i, name)
	}
}
