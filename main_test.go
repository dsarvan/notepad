package main

import (
	"testing"

	"fyne.io/fyne/test"
	"fyne.io/fyne/widget"
	"github.com/stretchr/testify/assert"
)

func testUIList(t *testing.T) {

	// running an application in memory
	gui := &ui{notes: []*note{
		{content: "1"},
		{content: "2"},
	}}

	// we don't care what is returning at this instance
	// because the gui object is holding refernce to the
	// things that we are going to test
	_ = gui.loadUI()

	// assert the number of items in the list should be two
	assert.Equal(t, 2, len(gui.list.Children))
}

func tapSetContent(t *testing.T) {

	// running an application in memory
	gui := &ui{notes: []*note{
		{content: "1"},
		{content: "2"},
	}}

	// we don't care what is returning at this instance
	// because the gui object is holding refernce to the
	// things that we are going to test
	_ = gui.loadUI()

	// assert content text == "1"
	assert.Equal(t, "1", gui.content.Text)

	// tap the next button [1] to the first button [0]
	test.Tap(gui.list.Children[1].(*widget.Button))
	// assert content text == "2"
	assert.Equal(t, "2", gui.content.Text)
}
