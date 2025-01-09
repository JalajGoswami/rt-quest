package main

import (
	"fmt"
	"log"
	"rt-quest/config"
	"rt-quest/luminous"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	sprites     []Sprite
	lightSource luminous.Illuminant
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
	g.lightSource.DrawRays(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.SCREEN_WIDTH, config.SCREEN_HEIGHT
}

func main() {
	var a = 2.8
	fmt.Println(int(a))
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("RT Demo")
	if err := ebiten.RunGame(&Game{
		sprites: []Sprite{
			// shape.NewCircle(50, 50, 40, color.White),
		},
		lightSource: luminous.NewIlluminant(20, 20),
	}); err != nil {
		log.Fatal(err)
	}
}
