package sprite

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite interface {
	Pos() (x, y float64)
	Draw(screen *ebiten.Image)
	Update() error
}
