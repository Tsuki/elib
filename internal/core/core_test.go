package core

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"strings"
	"testing"
)

func TestWalkDirectory(t *testing.T) {
	list, err := WalkDirectory("../../fixture", ".nfo")
	assert.NoError(t, err)
	assert.True(t, len(list) == 1)
	assert.True(t, strings.HasSuffix(list[0], "test.nfo"))
	t.Log(spew.Sdump(list))
}

func TestDecodeNfo(t *testing.T) {
	file, err := filepath.Abs("../../fixture/test.nfo")
	assert.NoError(t, err)
	got, err := DecodeNfo(file)
	assert.NoError(t, err)
	t.Log(spew.Sdump(got))
}
