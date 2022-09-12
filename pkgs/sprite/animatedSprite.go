package sprite

import (
	"image"

	"github.com/dragonchen-tw/cavestory-go/pkgs/graphics"
)

const (
	tileSize int = 32
)

type AnimatedSprite struct {
	Sprite
	frameTime     int // how many mileseconeds lasts per frame
	numFrames     int
	currentFrame  int   // which frame we currently on
	elapsedTime   int   // how many mileseconeds has elapsed since the last frame
	lastTimeStamp int64 // the last time (ms) we update the frame
}

func NewAnimatedSprite(
	gp graphics.Graphics, path string,
	left, top, width, height int,
	fps int, numFrames int,
) AnimatedSprite {
	return AnimatedSprite{
		Sprite:       NewSprite(gp, path, left, top, width, height),
		frameTime:    1000 / fps,
		numFrames:    numFrames,
		currentFrame: 0,
		elapsedTime:  0,
	}
}

func (as *AnimatedSprite) Update(elapsedGameTime int) {
	as.elapsedTime += elapsedGameTime

	if as.elapsedTime >= as.frameTime {
		as.elapsedTime = 0
		as.currentFrame++
		if as.currentFrame < as.numFrames {
			as.sourceRect = as.sourceRect.Add(image.Point{X: tileSize, Y: 0})
		} else {
			as.sourceRect = as.sourceRect.Sub(image.Point{X: tileSize * (as.numFrames - 1), Y: 0})
			as.currentFrame = 0
		}
	}
}

func (as *AnimatedSprite) Draw(gp graphics.Graphics, x int, y int) {
	as.Sprite.Draw(gp, x, y)
}
