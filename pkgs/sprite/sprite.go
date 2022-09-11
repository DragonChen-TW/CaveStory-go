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
	// Move the sprite to the center of the screen
	op.GeoM.Translate(
		float64(x)-float64(s.sourceReact.Dx()/2),
		float64(y)-float64(s.sourceReact.Dy()/2),
	)
	img := s.image.SubImage(s.sourceReact).(*ebiten.Image)
	gp.Blit(img, op)
}
