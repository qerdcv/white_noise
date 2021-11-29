package main

import (
	"fmt"
	"log"
	"white_noice/game"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	width  = 1000
	height = 700
)

func main() {
	g, err := game.New(width, height)

	if err != nil {
		log.Fatal(fmt.Errorf("failed to create game: %w", err))
	}

	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("White Noise")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
