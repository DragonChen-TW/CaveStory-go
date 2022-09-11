package sprite

import (
	"image"

	"github.com/dragonchen-tw/cavestory-go/pkgs/graphics"
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	image       *ebiten.Image
	sourceReact image.Rectangle
}

func NewSprite(
	gp graphics.Graphics, path string,
	left, top, width, height int,
) Sprite {
	img, ok := gp.LoadImage(path)
	if !ok {
		panic("Image loading is failed.")
	}
	return Sprite{
		image: img,
		sourceReact: image.Rect(
			left, top,
			left+width, top+height,
		),
	}
}

func (s Sprite) Blit(gp graphics.Graphics, x int, y int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	img := s.image.SubImage(s.sourceReact).(*ebiten.Image)
	gp.Blit(img, op)
}
