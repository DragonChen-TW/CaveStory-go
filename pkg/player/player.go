package player

import (
	"fmt"

	"github.com/dragonchen-tw/cavestory-go/pkg/graphics"
	"github.com/dragonchen-tw/cavestory-go/pkg/sprite"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Player struct {
	x, y   int
	sprite sprite.Spriter

	velX, velY int

	standing sprite.Sprite
	walking  sprite.AnimatedSprite
}

const (
	maxDash    = 0x45c
	maxMove    = 0x5ff
	dashGround = 0x60
	resist     = 0x33
	// max_dash: 0x32c,
	// max_move: 0x5ff,
	// gravity_air: 0x20,
	// gravity_ground: 0x50,
	// dash_air: 0x20,
	// dash_ground: 0x55,
	// resist: 0x33,
	// jump: 0x500,
)

func NewPlayer(gp graphics.Graphics, x, y int) Player {
	s := sprite.NewSprite(gp, "./imgs/chars.png",
		0, 0, 32, 32,
	)
	as := sprite.NewAnimatedSprite(gp, "./imgs/chars.png",
		0, 0, 32, 32,
		15, 3,
	)
	return Player{
		x:        x * 512,
		y:        y,
		sprite:   &s,
		standing: s,
		walking:  as,
	}
}

func (p *Player) Update(elapsedGameTime int) {
	// Debug
	// fmt.Println(p.x, p.y, p.velX)
	fmt.Println(ebiten.IsKeyPressed(ebiten.KeyLeft), inpututil.IsKeyJustPressed(ebiten.KeyLeft), inpututil.IsKeyJustReleased(ebiten.KeyLeft))

	p.sprite.Update(elapsedGameTime)

	// Movement (from doukutsu-rs)
	// Ref https://github.com/doukutsu-rs/doukutsu-rs/blob/a5f49c07e4a161f7ad5ecf97be7126db44f0b80e/src/game/player/mod.rs

	// ground
	onGround := p.y == 240
	if onGround {
		if inpututil.IsKeyJustPressed(ebiten.KeyLeft) || inpututil.IsKeyJustPressed(ebiten.KeyRight) {
			p.sprite = &p.walking
		}
		if inpututil.IsKeyJustReleased(ebiten.KeyLeft) || inpututil.IsKeyJustReleased(ebiten.KeyRight) {
			if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyRight) {
				p.sprite = &p.walking
			} else {
				p.sprite = &p.standing
			}
		}

		if ebiten.IsKeyPressed(ebiten.KeyLeft) && ebiten.IsKeyPressed(ebiten.KeyRight) {
			p.velX = 0
			p.sprite = &p.standing
		} else if ebiten.IsKeyPressed(ebiten.KeyLeft) && p.velX > -maxDash {
			p.velX -= dashGround
		} else if ebiten.IsKeyPressed(ebiten.KeyRight) && p.velX < maxDash {
			p.velX += dashGround
		}

		// Physics resist
		if p.velX < 0 {
			if p.velX > -resist {
				p.velX = 0
			} else {
				p.velX += resist
			}
		}
		if p.velX > 0 {
			if p.velX < resist {
				p.velX = 0
			} else {
				p.velX -= resist
			}
		}
	}

	// clamp vel_x between -max_move and max_move
	if p.velX < -maxMove {
		p.velX = -maxMove
	} else if p.velX > maxMove {
		p.velX = maxMove
	}

	if p.velX > resist || p.velX < -resist {
		p.x += p.velX
	}
}

func (p *Player) Draw(gp graphics.Graphics) {
	p.sprite.Draw(gp, p.x/512, p.y)
}
