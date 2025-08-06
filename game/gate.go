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

func (g *Gate) reset() {
	g.z = 15
	g.radius = 0
	g.angle = randomAngle()
	// g.angleIncrement = firefly.Degrees(float32(firefly.GetRandom()%200-100) * 0.05)
	g.angleIncrement = firefly.Degrees(0)
	g.passed = false
}

func (g *Gate) update() bool {
	const x = 0.0
	const d = 5.0
	const radius = 5.0
	g.z = g.z - 0.2
	xpLeft := (x - radius) * d / g.z
	xcLeft := xpLeft * 400.0 / 400.0
	xpRight := (x + radius) * d / g.z
	xcRight := xpRight * 400.0 / 400.0
	newRadius := xcRight - xcLeft
	g.radius = newRadius
	g.angle = firefly.Radians(g.angle.Radians() + g.angleIncrement.Radians())
	return newRadius <= firefly.Width/2
}

func (g *Gate) render() {
	color := firefly.ColorWhite
	if g.z > 10 {
		color = firefly.ColorLightGray
	}
	style := firefly.Outlined(color, int(g.radius*0.1))
	firefly.DrawArc(
		firefly.P(
			firefly.Width/2-int(g.radius),
			firefly.Height/2-int(g.radius),
		),
		int(g.radius)*2,
		g.angle,
		firefly.Radians(g.angle.Radians()+math.Pi/2*3),
		style,
	)
}

func randomAngle() firefly.Angle {
	return firefly.Degrees(float32(firefly.GetRandom() % 360))
}
