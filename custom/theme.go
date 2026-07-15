package custom

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type CustomTheme struct{}

func (t *CustomTheme) Color(
	colorName fyne.ThemeColorName,
	colorVariant fyne.ThemeVariant,
) color.Color {
	themeColor := theme.DefaultTheme().Color(colorName, colorVariant)

	primaryColor := color.RGBA{241, 81, 86, 255}
	secundaryColor := color.RGBA{107, 15, 26, 255}
	backgroundColor := color.RGBA{49, 8, 31, 255}

	disabledPrimaryColor := color.RGBA{165, 156, 156, 255}
	disabledSecundaryColor := color.RGBA{64, 58, 59, 255}

	/* switch colorVariant {
	case theme.VariantDark: */
	switch colorName {
	case theme.ColorNamePrimary:
		themeColor = primaryColor
	case theme.ColorNameForeground:
		themeColor = secundaryColor
	case theme.ColorNameBackground:
		themeColor = backgroundColor
	case theme.ColorNameDisabled:
		themeColor = disabledSecundaryColor
	case theme.ColorNameButton:
		themeColor = primaryColor
	case theme.ColorNameDisabledButton:
		themeColor = disabledPrimaryColor
	case theme.ColorNameError:
		themeColor = secundaryColor

	}

	// #f15156
	// #6b0f1a
	// #31081f

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
