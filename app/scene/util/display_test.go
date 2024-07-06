package util

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListOptions(t *testing.T) {
	buf := bytes.Buffer{}
	ctx := GameCtx{Writer: &buf}

	opts := []string{"zero", "one", "two"}

	ListOptions(ctx, opts)

	assert.Equal(t, "0: zero\n1: one\n2: two\n", buf.String())
}

func TestGameCtx_print(t *testing.T) {
	buf := bytes.Buffer{}
	ctx := GameCtx{Writer: &buf}

	ctx.Print("testing")

	assert.Equal(t, "testing", buf.String())
}

func TestGameCtx_printf(t *testing.T) {
	buf := bytes.Buffer{}
	ctx := GameCtx{Writer: &buf}

	ctx.Printf("%s is a %s", "truck", "car")

	assert.Equal(t, "truck is a car", buf.String())
}
