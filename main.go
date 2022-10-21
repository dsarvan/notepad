package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/widget"
)

func loadUI() fyne.CanvasObject {

	list := widget.NewVBox()
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {}),
	)

	titlelist := fyne.NewContainerWithLayout(Layout.NewBorderLayout(toolbar, nil, nil, nil), toolbar, list)
	notecontent := widget.NewMultiLineEntry()
	return widget.NewHSplitContainer(titlelist, notecontent)
}

func main() {
	a := app.New()
	w := a.NewWindow("Notepad")
	loadUI()
	w.ShowAndRun()
}
