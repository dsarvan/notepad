package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func loadUI() fyne.CanvasObject {

	list := widget.NewVBox(
		widget.NewLabel("Note 1"),
		widget.NewLabel("Note 2"),
	)
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {}),
	)

	titlelist := fyne.NewContainerWithLayout(layout.NewBorderLayout(toolbar, nil, nil, nil), toolbar, list)
	notecontent := widget.NewMultiLineEntry()

	split := widget.NewHSplitContainer(titlelist, notecontent)
	split.Offset = 0.25

	return split
}

func main() {
	a := app.New()
	w := a.NewWindow("Notepad")

	w.SetContent(loadUI())
	w.ShowAndRun()
}
