package shape

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Circle struct {
	x, y  float64
	rad   int
	color color.Color
}

func (c Circle) Draw(screen *ebiten.Image) {
	rad_squared := math.Pow(float64(c.rad), 2)
	diam := 2 * c.rad
	img := ebiten.NewImage(diam, diam)
	for x := range c.rad {
		// only considering positive cordinates as center is at origin
		// all quadrants will have same image
		for y := range c.rad {
			// check if x,y is within circumference
			isPointInCircle := true
			if !(x < c.rad/2 && y < c.rad/2) { // for optimization
				dist_squared := math.Pow(float64(x), 2) + math.Pow(float64(y), 2)
				if dist_squared > rad_squared {
					isPointInCircle = false
				}
			}

			if isPointInCircle { // for all 4 quadrants
				for _, sign := range [...]int{-1, 1} {
					img.Set(c.rad+sign*x, c.rad+sign*y, c.color)
					img.Set(c.rad-sign*x, c.rad+sign*y, c.color)
				}
			}
		}
	}
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(c.x, c.y)
	screen.DrawImage(img, opts)
}

func (c Circle) Update() error {
	return nil
}

func NewCircle(x, y float64, rad int, c color.Color) Circle {
	return Circle{x, y, rad, c}
}
