package ui

import (
	"fyne.io/fyne/v2"
	"github.com/vitorlfaria/pixl/internal/apptype"
	"github.com/vitorlfaria/pixl/internal/pxcanvas"
	"github.com/vitorlfaria/pixl/internal/swatch"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
