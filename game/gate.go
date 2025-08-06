package game

import (
	"math"

	"github.com/firefly-zero/firefly-go/firefly"
)

type Gate struct {
	z              float32
	radius         float32
	angle          firefly.Angle
	angleIncrement firefly.Angle
	passed         bool
}

func newGate(count int) *Gate {
	return &Gate{
		z:              30,
		radius:         0,
		angle:          randomAngle(),
		angleIncrement: firefly.Degrees(float32(randomInt(-100, 100)) * 0.04),
		passed:         false,
	}
}

func (g *Gate) update() bool {
	const x = 0.0
	const d = 8.0
	const radius = 8.0
	g.z = g.z - 0.2
	xLeft := (x - radius) * d / g.z
	xRight := (x + radius) * d / g.z
	newRadius := xRight - xLeft
	g.radius = newRadius
	newAngle := g.angle.Radians() + g.angleIncrement.Radians()
	if newAngle > math.Pi*2 {
		newAngle -= math.Pi * 2
	} else if newAngle < 0 {
		newAngle += math.Pi * 2
	}
	g.angle = firefly.Radians(newAngle)
	return newRadius <= firefly.Width/2
}

func (g *Gate) render() {
	color := firefly.ColorWhite
	if g.z > 15 {
		color = firefly.ColorLightGray
	}
	width := max(1, int(g.radius*0.1))
	style := firefly.Outlined(color, width)
	firefly.DrawArc(
		firefly.P(
			firefly.Width/2-int(g.radius),
			firefly.Height/2-int(g.radius),
		),
		int(g.radius*2),
		g.angle,
		firefly.Radians(math.Pi*3/2),
		style,
	)
}
