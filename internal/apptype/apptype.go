package apptype

import (
	"image/color"

	"fyne.io/fyne/v2"
)

type BushType = int

type PxCanvasConfig struct {
	DrawingArea    fyne.Size
	CanvasOffset   fyne.Position
	PxRows, PxCols int
	PXSize         int
}

type State struct {
	BrushColor     color.Color
	BrushType      BushType
	SwatchSelected int
	FilePath       string
}

func (statue *State) SetFilePath(path string) {
	statue.FilePath = path
}
