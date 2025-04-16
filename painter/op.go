package painter

import (
	"image"
	"image/color"

	"golang.org/x/exp/shiny/screen"
	"golang.org/x/image/draw"
)


type Operation interface {
	Do(t screen.Texture) (ready bool)
}

type OperationList []Operation

func (ol OperationList) Do(t screen.Texture) (ready bool) {
	for _, o := range ol {
		ready = o.Do(t) || ready
	}
	return
}


var UpdateOp = updateOp{}

type updateOp struct{}

func (op updateOp) Do(t screen.Texture) bool { return true }


type OperationFunc func(t screen.Texture)

func (f OperationFunc) Do(t screen.Texture) bool {
	f(t)
	return false
}

func WhiteFill(t screen.Texture) {
	t.Fill(t.Bounds(), color.White, screen.Src)
}

func GreenFill(t screen.Texture) {
	t.Fill(t.Bounds(), color.RGBA{G: 0xff, A: 0xff}, screen.Src)
}

type BgRectangle struct {
	X1, Y1, X2, Y2 float64
}

func (op *BgRectangle) Do(t screen.Texture) bool {
	bounds := t.Bounds()
	width := float64(bounds.Dx())
	height := float64(bounds.Dy())
	
	t.Fill(image.Rect(
		int(op.X1 * width), int(op.Y1 * height),
		int(op.X2 * width), int(op.Y2 * height)),
		color.Black, screen.Src,
	)
	return false
}

type Figure struct {
	X, Y float64
	C    color.RGBA
}

func (op *Figure) Do(t screen.Texture) bool {
	bounds := t.Bounds()
	width := float64(bounds.Dx())
	height := float64(bounds.Dy())
	
	x := int(op.X * width)
	y := int(op.Y * height)
	
	barWidth := 60
	barLength := 200
	
	// Draw vertical bar (main part of T)
	t.Fill(image.Rect(
		x-barWidth/2,
		y-barLength/2,
		x+barWidth/2,
		y+barLength/2,
	), color.RGBA{255, 255, 0, 255}, draw.Src)
	
	// Draw horizontal bar (right part of T)
	t.Fill(image.Rect(
		x,
		y-barWidth/2,
		x+barLength/2,
		y+barWidth/2,
	), color.RGBA{255, 255, 0, 255}, draw.Src)
	
	return false
}

type Move struct {
	X, Y    float64
	Figures []*Figure
}

func (op *Move) Do(t screen.Texture) bool {
	for i := range op.Figures {
		op.Figures[i].X += op.X
		op.Figures[i].Y += op.Y
	}
	return false
}

func ResetScreen(t screen.Texture) {
	t.Fill(t.Bounds(), color.Black, draw.Src)
}