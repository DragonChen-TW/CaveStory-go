package graphics

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Graphics struct {
	Screen   *ebiten.Image
	imgCache map[string]*ebiten.Image
}

func NewGraphics() Graphics {
	gp := Graphics{
		Screen:   ebiten.NewImage(640, 480),
		imgCache: make(map[string]*ebiten.Image),
	}
	gp.Reset()
	return gp
}

func (g *Graphics) Reset() {
	g.Screen.Fill(color.RGBA{0, 0, 0, 255})
}

func (gp Graphics) LoadImage(path string) (img *ebiten.Image, ok bool) {
	if img, ok = gp.imgCache[path]; !ok {
		// Load image from file
		var err error
		img, _, err = ebitenutil.NewImageFromFile(path)
		if err != nil {
			panic(err)
		}

		// Cache the image
		gp.imgCache[path] = img
		ok = true
	}
	return img, ok
}

func (gp Graphics) Blit(src *ebiten.Image, op *ebiten.DrawImageOptions) {
	gp.Reset()
	gp.Screen.DrawImage(src, op)
}
