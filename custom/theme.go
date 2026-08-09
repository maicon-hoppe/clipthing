package custom

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

var (
	PrimaryColor    = color.RGBA{241, 81, 86, 255} 	// #f15156
	SecondaryColor  = color.RGBA{107, 15, 26, 255}	// #6b0f1a
	BackgroundColor = color.RGBA{49, 8, 31, 255}	// #31081f

	DisabledPrimaryColor   = color.RGBA{165, 156, 156, 255}
	DisabledSecondaryColor = color.RGBA{64, 58, 59, 255}

	TextColor = color.RGBA{204, 204, 204, 255}
)

type CustomTheme struct{}

func (t *CustomTheme) Color(
	colorName fyne.ThemeColorName,
	colorVariant fyne.ThemeVariant,
) color.Color {
	themeColor := theme.DefaultTheme().Color(colorName, colorVariant)

	/* switch colorVariant {
	case theme.VariantDark: */
	switch colorName {
	case theme.ColorNamePrimary:
		themeColor = PrimaryColor
	case theme.ColorNameForeground:
		themeColor = SecondaryColor
	case theme.ColorNameBackground:
		themeColor = BackgroundColor
	case theme.ColorNameDisabled:
		themeColor = DisabledSecondaryColor
	case theme.ColorNameButton:
		themeColor = PrimaryColor
	case theme.ColorNameDisabledButton:
		themeColor = DisabledPrimaryColor
	case theme.ColorNameError:
		themeColor = SecondaryColor

	}

	/* case theme.VariantLight:
		switch colorName {
		case theme.ColorNamePrimary:
			themeColor = color.RGBA{241, 84, 86, 255}

		}
	}*/

	return themeColor
}

func (t *CustomTheme) Font(fontName fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(fontName)
}

func (t *CustomTheme) Icon(iconName fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(iconName)
}

func (t *CustomTheme) Size(sizeName fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(sizeName)
}
