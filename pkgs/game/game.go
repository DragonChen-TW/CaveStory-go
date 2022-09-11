package game

import (
	"fmt"

	"github.com/dragonchen-tw/cavestory-go/pkgs/graphics"
	"github.com/dragonchen-tw/cavestory-go/pkgs/sprite"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	keys     []ebiten.Key
	graphics graphics.Graphics
	player   sprite.Sprite
}

func NewGame() Game {
	gp := graphics.NewGraphics()
	player := sprite.NewSprite(gp, "imgs/chars.png",
		0, 0, 32, 32,
	)

	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeOnlyFullscreenEnabled)
	return Game{
		keys:     make([]ebiten.Key, 0),
		graphics: gp,
		player:   player,
	}
}

func (g *Game) Update() error {
	// Test Sprite
	w, h := g.graphics.Screen.Size()
	g.player.Blit(g.graphics, w/2, h/2)

	// Keyboard
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	for _, key := range g.keys {
		if key == ebiten.KeyEscape {
			// close the game loop
			return fmt.Errorf("Exit Game")
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.graphics.Screen, nil)
	ebitenutil.DebugPrint(screen, "Hello, World!")
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("FPS: %.02f", ebiten.CurrentTPS()), 0, 12)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
