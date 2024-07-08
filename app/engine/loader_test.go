package engine_test

import (
	"embed"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

//go:embed asset/*
var asset embed.FS

func TestLoadAssets(t *testing.T) {
	dir, err := asset.ReadDir("asset")
	if err != nil {
		panic(err)
	}

	for _, o := range dir {
		fmt.Println(!o.IsDir(), ":", o.Name())
	}
	assert.Fail(t, "It's gonna happen")
}
