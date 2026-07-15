package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"golang.design/x/clipboard"

	"github.com/maicon-hoppe/clipthing/custom"
)

func main() {
	myApp := app.NewWithID("com.clipthing")
	myApp.Settings().SetTheme(&custom.CustomTheme{})

	err := clipboard.Init()
	if err != nil {
		fmt.Println("An error occurred: ")
		fmt.Println(err)
		return
	}

	w := myApp.NewWindow("Teste")
	w.Resize(fyne.NewSize(200, 300))

	clipboardText := string(clipboard.Read(clipboard.FmtText))
	w.SetContent(custom.NewClipboardItemWidget(clipboardText, custom.TextDisplay))
	w.ShowAndRun()
}
