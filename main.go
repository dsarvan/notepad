package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

type ui struct {
	current *note   // current note
	notes   []*note // notes list

	notecontent *widget.Entry // user interface components
	list        *widget.Box   // highlight the selected label
}

func (u *ui) setNote(n *note) {
	u.notecontent.SetText(n.content)
	u.current = n
	u.refreshList(u.notes)
}

func (u *ui) refreshList(notes []*note) {
	u.list.Children = nil
	for _, n := range u.notes {
		thisNote := n

		notelabel := widget.NewButton(n.title(), func() {
			u.setNote(thisNote)
		})

		// highlight the label of the selected note
		if n == u.current {
			notelabel.Style = widget.PrimaryButton
		}

		u.list.Append(notelabel)
	}
}

func (u *ui) loadUI() fyne.CanvasObject {

	u.notecontent = widget.NewMultiLineEntry()

	// highlight the selected label
	u.list = widget.NewVBox()
	u.refreshList(u.notes)

	// show the Note 1 content by default
	if len(u.notes) > 0 {
		u.setNote(u.notes[0])
	}

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {}),
	)

	titlelist := fyne.NewContainerWithLayout(layout.NewBorderLayout(toolbar, nil, nil, nil), toolbar, u.list)

	split := widget.NewHSplitContainer(titlelist, u.notecontent)
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

	notesUI := &ui{notes: notelist}
	w.SetContent(notesUI.loadUI())
	w.ShowAndRun()
}
