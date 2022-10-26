package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testNoteTitle(t *testing.T) {
	tnote := &note{content: "Note Title"}
	assert.Equal(t, "Note Title", tnote.title())

	// new instance of note to return content in line 1
	// let first line of the text be called title
	tnote = &note{content: "Note\nLine2"}
	assert.Equal(t, "Note", tnote.title())

	// no content on the note
	tnote = &note{content: " "}
	assert.Equal(t, "Untitled", tnote.title())
}
