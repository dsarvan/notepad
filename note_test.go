package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testNoteTitle(t *testing.T) {
	tnote := &note{content: "Note Title"}
	assert.Equal(t, "Note Title", tnote.title())
}
