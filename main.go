package main

import (
	"log"

	g "github.com/dragonchen-tw/cavestory-go/pkg/game"

	"github.com/hajimehoshi/ebiten/v2"
)

// Author:	DragonChen https://github.com/dragonchen-tw/
// Title:	Main File for test
// Date:	2022/09/10

func main() {
	ebiten.SetWindowSize(640, 480)
	game := g.NewGame()
	if err := ebiten.RunGame(&game); err != nil {
		if err.Error() == "Exit Game" {
			log.Println("Exit Game")
		} else {
			log.Fatal(err)
		}
	}
}
