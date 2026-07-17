package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
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
	w.Resize(fyne.NewSize(1000, 600))

	clipboardText := string(clipboard.Read(clipboard.FmtText))
	clipboardItems := container.New(&custom.VFlex{}, custom.NewClipboardItemWidget(clipboardText, custom.TextDisplay))

	// ch := clipboard.Watch(context.Background())
	/*go func() {
		for item := range ch {
			switch item.Format {
			case clipboard.FmtText:
				fyne.Do(func() {
					clipboardItem := custom.NewClipboardItemWidget(string(item.Bytes), custom.TextDisplay)
					clipboardItems.Add(clipboardItem)
				})
			case clipboard.FmtImage:
				fyne.Do(func() {
					imageItem := custom.NewClipboardItemWidget("Image", custom.TextDisplay)
					clipboardItems.Add(imageItem)
				})
			}
		}
	}() */

	clipboardItems.Add(custom.NewClipboardItemWidget("outra nota", custom.TextDisplay))
	clipboardItems.Add(custom.NewClipboardItemWidget("{308 108}", custom.TextDisplay))

	w.SetContent(clipboardItems)
	w.ShowAndRun()
}
