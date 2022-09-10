package game

import (
	"fmt"

	"github.com/dragonchen-tw/cavestory-go/pkgs/graphics"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	keys     []ebiten.Key
	graphics graphics.Graphics
}

func NewGame() Game {
	return Game{
		keys:     make([]ebiten.Key, 0),
		graphics: graphics.NewGraphics(),
	}
}

func (g *Game) Update() error {
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
	g.graphics.Draw(screen)
	ebitenutil.DebugPrint(screen, "Hello, World!")
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("FPS: %.02f", ebiten.CurrentTPS()), 0, 12)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
