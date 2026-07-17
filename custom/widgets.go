package custom

import (
	"golang.org/x/image/colornames"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type DisplayMode int

const (
	TextDisplay DisplayMode = iota
	ImageDisplay
)

type ClipboardItem struct {
	widget.BaseWidget
	content string
	mode    DisplayMode
}

func (item *ClipboardItem) CreateRenderer() fyne.WidgetRenderer {
	bg := canvas.NewRectangle(theme.Color(theme.ColorNamePrimary))
	bg.SetMinSize(fyne.NewSize(300, 100))
	bg.CornerRadius = theme.Size(theme.SizeNameCardRadius)

	card := container.NewPadded(bg)
	cardView := container.NewVBox()

	if item.mode == TextDisplay {
		content := widget.NewLabel(item.content)
		content.Wrapping = fyne.TextWrapWord
		cardView.Add(content)
		cardView.Add(layout.NewSpacer())
	}

	cardView.Add(widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {
			bg.FillColor = colornames.Blueviolet
		}),
	))
	card.Add(cardView)

	return widget.NewSimpleRenderer(card)
}

func NewClipboardItemWidget(content string, mode DisplayMode) *ClipboardItem {
	clipItem := &ClipboardItem{
		content: content,
		mode:    mode,
	}
	clipItem.ExtendBaseWidget(clipItem)

	return clipItem
}
