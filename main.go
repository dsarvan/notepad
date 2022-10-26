package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func loadUI(notes []*note) fyne.CanvasObject {

	notecontent := widget.NewMultiLineEntry()

	list := widget.NewVBox()
	for _, n := range notes {
		thisNote := n
		list.Append(widget.NewButton(n.title(), func() {
			notecontent.SetText(thisNote.content)
		}))
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
