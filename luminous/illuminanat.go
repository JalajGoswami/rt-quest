package luminous

import (
	"image/color"
	"log"
	"math"
	"rt-quest/config"
	"rt-quest/sprite"

	"github.com/hajimehoshi/ebiten/v2"
)

type Illuminant struct {
	x, y         float64
	rayCount     int
	sourceObject sprite.Sprite
}

func NewIlluminant(s sprite.Sprite, rayCount int) Illuminant {
	x, y := s.Pos()
	return Illuminant{x, y, rayCount, s}
}

func (i *Illuminant) DrawRays(screen *ebiten.Image) {
	if i.rayCount == 0 {
		log.Fatal("Ray count cannot be zero")
	}
	delta := 2 * math.Pi / float64(i.rayCount)
	angle := 0.0
	for range i.rayCount {
		for x := range config.SCREEN_WIDTH {
			y := math.Tan(angle)*(float64(x)-i.x) + i.y
			if y <= config.SCREEN_HEIGHT && y >= 0 {
				screen.Set(x, int(y), color.RGBA{255, 255, 0, 255})
			}
		}
		angle += delta
	}
}

func (i *Illuminant) Update() {
	i.x, i.y = i.sourceObject.Pos()
}
