package sprite

import (
	"image"

	"github.com/dragonchen-tw/cavestory-go/pkg/graphics"
	"github.com/hajimehoshi/ebiten/v2"
)

type Spriter interface {
	Update(int) // update the sprite
	Draw(gp graphics.Graphics, x int, y int)
}

type Sprite struct {
	image      *ebiten.Image
	sourceRect image.Rectangle
}

func NewSprite(
	gp graphics.Graphics,
	path string,
	left, top, width, height int,
) Sprite {
	img, ok := gp.LoadImage(path)
	if !ok {
		panic("Image loading is failed.")
	}
	return Sprite{
		image: img,
		sourceRect: image.Rect(
			left, top,
			left+width, top+height,
		),
	}
}

func (s *Sprite) Update(int) {}

func (s *Sprite) Draw(gp graphics.Graphics, x int, y int) {
	op := &ebiten.DrawImageOptions{}
	// Move the sprite to the center of the screen
	op.GeoM.Translate(
		float64(x)-float64(s.sourceRect.Dx()/2),
		float64(y)-float64(s.sourceRect.Dy()/2),
	)
	img := s.image.SubImage(s.sourceRect).(*ebiten.Image)
	gp.Blit(img, op)
}
