package custom

import (
	"math"
	"slices"

	"fyne.io/fyne/v2"
)

type VFlex struct {
	ContainerSize fyne.Size
}

func (l *VFlex) MinSize(objects []fyne.CanvasObject) fyne.Size {
	if len(objects) == 0 {
		return fyne.NewSize(0, 0)
	}

	minimumHeights := make([]float32, len(objects))
	minimumWidths := make([]float32, len(objects))
	for _, obj := range objects {
		minimumHeights = append(minimumHeights, obj.MinSize().Height)
		minimumWidths = append(minimumWidths, obj.MinSize().Width)
	}
	maxSmallHeight := slices.Max(minimumHeights)
	maxSmallWidth := slices.Max(minimumWidths)

	minimumHeight := maxSmallHeight
	maxSmallHeightIndex := slices.IndexFunc(objects,
		func(obj fyne.CanvasObject) bool {
			return obj.MinSize().Height == maxSmallHeight
		},
	)
	squareProportion := math.Sqrt(float64(len(objects)))
	if squareProportion > 1 {
		step := int(math.Ceil(squareProportion))
		for i := maxSmallHeightIndex - step; i >= 0; i -= step {
			minimumHeight += objects[i].MinSize().Height
		}

		for i := maxSmallHeightIndex + step; i <= (len(objects) - 1); i += step {
			minimumHeight += objects[i].MinSize().Height
		}
	}

	return fyne.NewSize(maxSmallWidth, minimumHeight)
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
