package engine_test

import (
	"embed"
	asset2 "github.com/coopstools/minibeast/app/asset"
	"github.com/stretchr/testify/assert"
	"testing"
)

//go:embed asset/*
var asset embed.FS

func TestLoadAssets(t *testing.T) {
	asset2.LoadFont()
	assert.Fail(t, "It's gonna happen")
}
