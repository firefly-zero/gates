package game

import (
	"github.com/firefly-zero/firefly-go/firefly"
	"github.com/orsinium-labs/tinymath"
)

type TitleGate struct {
	angle firefly.Angle
	scale float32
}

func (g *TitleGate) render() {
	r := int(40 * g.scale)
	width := max(1, int(2*g.scale))
	firefly.DrawArc(
		firefly.P(
			firefly.Width/2-r,
			firefly.Height/2-r,
		),
		r*2,
		g.angle,
		firefly.Radians(tinymath.Pi*3/2),
		firefly.Outlined(firefly.ColorWhite, width),
	)
}
