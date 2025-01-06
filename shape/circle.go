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

func NewCircle(x, y float64, rad int, c color.Color) *Circle {
	return &Circle{x, y, rad, c}
}

func (c *Circle) Draw(screen *ebiten.Image) {
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
	opts.GeoM.Translate(c.x-float64(c.rad), c.y-float64(c.rad))
	screen.DrawImage(img, opts)
}

func (c *Circle) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if c.isPointInCircle(float64(x), float64(y)) {
			c.x, c.y = float64(x), float64(y)
		}
	}
	return nil
}

func (c *Circle) isPointInCircle(x, y float64) bool {
	if x > c.x+float64(c.rad) || x < c.x-float64(c.rad) ||
		y > c.y+float64(c.rad) || y < c.y-float64(c.rad) {
		return false
	}

	rad_squared := math.Pow(float64(c.rad), 2)
	dist_squared := math.Pow(x-c.x, 2) + math.Pow(y-c.y, 2)
	if dist_squared > rad_squared {
		return false
	}
	return true
}
