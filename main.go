package main

import (
	"context"
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	_ "image/gif"
	_ "image/jpeg"

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
	windowSize := fyne.NewSize(1000, 600)
	w.Resize(windowSize)

	clipboardText := clipboard.Read(clipboard.FmtText)
	clipboardItems := container.New(
		&custom.VFlex{ContainerSize: windowSize},
		custom.NewClipboardItemWidget(clipboardText, custom.TextDisplay),
	)
	scrollableContainer := container.NewVScroll(clipboardItems)

	ch := clipboard.Watch(context.Background())
	go func() {
		for item := range ch {
			switch item.Format {
			case clipboard.FmtText:
				fyne.Do(func() {
					clipboardItem := custom.NewClipboardItemWidget(item.Bytes, custom.TextDisplay)
					clipboardItems.Add(clipboardItem)
					clipboardItems.Refresh()
				})
			case clipboard.FmtImage:
				_, _, err := image.Decode(bytes.NewReader(item.Bytes))
				if err == nil {
					fyne.Do(func() {
						imageItem := custom.NewClipboardItemWidget(item.Bytes, custom.ImageDisplay)
						clipboardItems.Add(imageItem)
						clipboardItems.Refresh()
					})
				}
			}
		}
	}() 

	w.SetContent(scrollableContainer)
	w.ShowAndRun()
}
