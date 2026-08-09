package main

import (
	"bytes"
	"context"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"golang.design/x/clipboard"
	"golang.org/x/image/colornames"

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

	w := myApp.NewWindow("Clipthing")
	windowSize := fyne.NewSize(1000, 600)
	w.Resize(windowSize)

	if desk, ok := myApp.(desktop.App); ok {
		isWindowHidden := false
		w.SetCloseIntercept(func() {
			w.Hide()
			isWindowHidden = true
		})

		menu := fyne.NewMenu("Clipthing",
			fyne.NewMenuItem("Show/Hide", func() {
				if isWindowHidden {
					w.Show()
					isWindowHidden = false
				} else {
					w.Hide()
					isWindowHidden = true
				}
			}),
		)
		desk.SetSystemTrayMenu(menu)
	}

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
				contentIsValidSVG := strings.Contains(string(item.Bytes), "<svg ") &&
					strings.Contains(string(item.Bytes), "</svg>") &&
					len(item.Bytes) > 11

				if contentIsValidSVG {
					fyne.Do(func() {
						imageItem := custom.NewClipboardItemWidget(item.Bytes, custom.ImageDisplay)
						clipboardItems.Add(imageItem)
						clipboardItems.Refresh()
					})
				} else {
					fyne.Do(func() {
						clipboardItem := custom.NewClipboardItemWidget(item.Bytes, custom.TextDisplay)
						clipboardItems.Add(clipboardItem)
						clipboardItems.Refresh()
					})
				}
			case clipboard.FmtImage:
				_, _, err := image.Decode(bytes.NewReader(item.Bytes))
				if err != nil {
					continue
				} else {
					fyne.Do(func() {
						imageItem := custom.NewClipboardItemWidget(item.Bytes, custom.ImageDisplay)
						clipboardItems.Add(imageItem)
						clipboardItems.Refresh()
					})
				}
			}
		}
	}()

	iconWidget := canvas.NewImageFromResource(myApp.Icon())
	iconWidget.FillMode = canvas.ImageFillContain
	iconWidget.SetMinSize(fyne.NewSquareSize(30))
	title := canvas.NewText("CLIPTHING", colornames.White)

	title.FontSource = custom.ResourcePTSansBoldTtf
	header := container.NewVBox(
		container.NewHBox(
			iconWidget,
			title,
		),
		canvas.NewLine(custom.PrimaryColor),
	)
	globalContainer := container.NewBorder(header, nil, nil, nil, scrollableContainer)
	w.SetContent(globalContainer)
	w.ShowAndRun()
}
