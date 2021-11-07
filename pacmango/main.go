package main

import (
	"log"

	_ "image/png"

	"github.com/hajimehoshi/ebiten"
	"github.com/kgosse/pacmango/pacman"
)

func main() {
	g := pacman.NewGame()

	if err := ebiten.Run(g.Update, g.ScreenWidth(), g.ScreenHeight(), 1, "Pacman"); err != nil {
		log.Fatal(err)
	}
}
