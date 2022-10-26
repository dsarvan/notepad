package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

var (
	current     *note
	notecontent *widget.Entry
)

func setNote(n *note) {
	notecontent.SetText(n.content)
	current = n
}

func loadUI(notes []*note) fyne.CanvasObject {

	notecontent = widget.NewMultiLineEntry()

	// show the Note 1 content by default
	if len(notes) > 0 {
		setNote(notes[0])
	}

	list := widget.NewVBox()
	for _, n := range notes {
		thisNote := n

		notelabel := widget.NewButton(n.title(), func() {
			setNote(thisNote)
		})

		// highlight the label of the selected note
		if n == current {
			notelabel.Style = widget.PrimaryButton
		}

		list.Append(notelabel)
	}

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {}),
	)

	titlelist := fyne.NewContainerWithLayout(layout.NewBorderLayout(toolbar, nil, nil, nil), toolbar, list)

	split := widget.NewHSplitContainer(titlelist, notecontent)
	split.Offset = 0.25

	return split
}

func main() {
	a := app.New()
	w := a.NewWindow("Notepad")

	// list of notes
	notelist := []*note{
		{content: "Note 1\nContent 1"},
		{content: "Note 2\nContent 2"},
	}

	w.SetContent(loadUI(notelist))
	w.ShowAndRun()
}
