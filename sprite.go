package main

import "github.com/hajimehoshi/ebiten/v2"

type Sprite interface {
	Draw(screen *ebiten.Image)
	Update() error
}
