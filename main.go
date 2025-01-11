package main

import (
	"image/color"
	"log"
	"rt-quest/config"
	"rt-quest/luminous"
	"rt-quest/shape"
	"rt-quest/sprite"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	sprites     []sprite.Sprite
	lightSource luminous.Illuminant
}

func (g *Game) Update() error {
	for _, sprite := range g.sprites {
		if err := sprite.Update(); err != nil {
			return err
		}
	}
	g.lightSource.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.lightSource.DrawRays(screen)
	for _, sprite := range g.sprites {
		sprite.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.SCREEN_WIDTH, config.SCREEN_HEIGHT
}

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("RT Demo")

	s := shape.NewCircle(50, 50, 40, color.White)
	illuminant := luminous.NewIlluminant(s, 51)

	if err := ebiten.RunGame(&Game{
		sprites:     []sprite.Sprite{s},
		lightSource: illuminant,
	}); err != nil {
		log.Fatal(err)
	}
}
