package ui

import (
	"fyne.io/fyne/v2"
	"github.com/vitorlfaria/pixl/internal/apptype"
	"github.com/vitorlfaria/pixl/internal/swatch"
)

type AppInit struct {
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
