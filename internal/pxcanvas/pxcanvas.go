package pxcanvas

import (
	"image"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/vitorlfaria/pixl/internal/apptype"
)

type PxCanvasMouseState struct {
	previousCoord *fyne.PointEvent
}

type PxCanvas struct {
	widget.BaseWidget
	apptype.PxCanvasConfig
	renderer     *PxCanvasRenderer
	PixelData    image.Image
	mouseState   PxCanvasMouseState
	appState     *apptype.State
	reloadImgage bool
}

func (pxCanvas *PxCanvas) Bounds() image.Rectangle {
	x0 := int(pxCanvas.CanvasOffset.X)
	y0 := int(pxCanvas.CanvasOffset.Y)
	x1 := int(pxCanvas.PxCols*pxCanvas.PXSize + int(pxCanvas.CanvasOffset.X))
	y1 := int(pxCanvas.PxRows*pxCanvas.PXSize + int(pxCanvas.CanvasOffset.Y))
	return image.Rect(x0, y0, x1, y1)
}

func InBounds(pos fyne.Position, bounds image.Rectangle) bool {
	x := int(pos.X)
	y := int(pos.Y)
	return x >= bounds.Min.X && x < bounds.Max.X && y >= bounds.Min.Y && y < bounds.Max.Y
}

func NewBlankImage(cols, rows int, c color.Color) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, cols, rows))
	for x := 0; x < cols; x++ {
		for y := 0; y < rows; y++ {
			img.Set(x, y, c)
		}
	}
	return img
}

func NewPxCanvas(config apptype.PxCanvasConfig, appState *apptype.State) *PxCanvas {
	pxCanvas := &PxCanvas{
		PxCanvasConfig: config,
		appState:       appState,
	}
	pxCanvas.PixelData = NewBlankImage(config.PxCols, config.PxRows, color.NRGBA{128, 128, 128, 255})
	pxCanvas.ExtendBaseWidget(pxCanvas)
	return pxCanvas
}

func (pxCanvas *PxCanvas) CreateRenderer() fyne.WidgetRenderer {
	canvasImage := canvas.NewImageFromImage(pxCanvas.PixelData)
	canvasImage.ScaleMode = canvas.ImageScalePixels
	canvasImage.FillMode = canvas.ImageFillContain

	canvasBorder := make([]canvas.Line, 4)
	for i := 0; i < len(canvasBorder); i++ {
		canvasBorder[i] = canvas.Line{StrokeColor: color.NRGBA{100, 100, 100, 255}, StrokeWidth: 2}
	}

	renderer := &PxCanvasRenderer{
		canvasImage:  canvasImage,
		canvasBorder: canvasBorder,
		pxCanvas:     pxCanvas,
	}
	pxCanvas.renderer = renderer
	return renderer
}
