package custom

import (
	"math"
	"slices"

	"fyne.io/fyne/v2"
)

type VFlex struct{}

func (l *VFlex) MinSize(objects []fyne.CanvasObject) fyne.Size {
	totalHeight := float32(0)
	minimumWidths := make([]float32, len(objects))
	for _, obj := range objects {
		totalHeight += obj.MinSize().Height
		minimumWidths = append(minimumWidths, obj.MinSize().Width)
	}
	maxSmallWidth := slices.Max(minimumWidths)

	return fyne.NewSize(maxSmallWidth, totalHeight)
}

func (l *VFlex) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	if len(objects) > 0 && !containerSize.IsZero() {
		minimumWidths := make([]float32, len(objects))
		for _, obj := range objects {
			obj.Resize(obj.MinSize())
			minimumWidths = append(minimumWidths, obj.MinSize().Width)
		}
		maxSmallWidth := slices.Max(minimumWidths)
		itemsPerRow := int(math.Floor(float64(containerSize.Width / maxSmallWidth)))

		objectsPerColumn := make([][]int, len(objects))
		for idx := 0; idx < len(objects); idx++ {
			obj := objects[idx]

			column := idx % itemsPerRow
			columnHeight := float32(0)
			for _, columnObjectIdx := range objectsPerColumn[column] {
				columnHeight += objects[columnObjectIdx].Size().Height
			}

			pos := fyne.NewPos(
				obj.Size().Width*float32(column),
				columnHeight,
			)
			obj.Move(pos)

			objectsPerColumn[column] = append(objectsPerColumn[column], idx)
		}
	}
}
