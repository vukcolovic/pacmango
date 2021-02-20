package main

import (
	"github.com/vukcolovic/pacmango/pacman"
	"log"
	"github.com/hajimehoshi/ebiten"
)

func main() {
	g := pacman.NewGame()

	if err := ebiten.Run(g.Update, g.ScreenWidth(), g.ScreenHeight(), 2, "Pacman"); err != nil {
		log.Fatal(err)
	}
}