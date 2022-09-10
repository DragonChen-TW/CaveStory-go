package graphics

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Graphics struct {
	screen   *ebiten.Image
	imgCache map[string]*ebiten.Image
}

func NewGraphics() Graphics {
	g := Graphics{
		screen:   ebiten.NewImage(640, 480),
		imgCache: make(map[string]*ebiten.Image),
	}
	g.screen.Fill(color.RGBA{0x30, 0x30, 0x30, 0xff})
	return g
}

func (g Graphics) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.screen, nil)
}
