package main

import (
	"image/color"
	"log"
	"rt-demo/config"
	"rt-demo/shape"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	sprites []Sprite
}

func (g *Game) Update() error {
	for _, sprite := range g.sprites {
		if err := sprite.Update(); err != nil {
			return err
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, sprite := range g.sprites {
		sprite.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.SCREEN_WIDTH, config.SCREEN_HEIGHT
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("RT Demo")
	if err := ebiten.RunGame(&Game{
		sprites: []Sprite{
			shape.NewCircle(50, 50, 50, color.White),
		},
	}); err != nil {
		log.Fatal(err)
	}
}
