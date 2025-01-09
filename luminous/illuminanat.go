package luminous

import (
	"image/color"
	"log"
	"math"
	"rt-quest/config"

	"github.com/hajimehoshi/ebiten/v2"
)

type Illuminant struct {
	x, y     float64
	rayCount int
}

func NewIlluminant(x, y float64) Illuminant {
	return Illuminant{x, y, 3}
}

var first = true

func (i *Illuminant) DrawRays(screen *ebiten.Image) {
	if i.rayCount == 0 {
		log.Fatal("Ray count cannot be zero")
	}
	delta := 2 * math.Pi / float64(i.rayCount)
	angle := 0.0
	for range i.rayCount {
		for x := range config.SCREEN_WIDTH {
			y := math.Tan(angle)*(float64(x)-i.x) + i.y
			if y > config.SCREEN_HEIGHT || y < 0 {
				break
			}
			screen.Set(x, int(y), color.White)
		}
		angle += delta
	}
}
