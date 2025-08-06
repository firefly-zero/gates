package game

import (
	"math"

	"github.com/firefly-zero/firefly-go/firefly"
)

type Gate struct {
	z        float32
	radius   float32
	angle    firefly.Angle
	angleInc firefly.Angle
	passed   bool
}

func newGate(count int) *Gate {
	angleInc := float32(randomInt(-100, 100)) * getAngleInc(count)
	angle := float32(getAngle(count))
	return &Gate{
		z:        30,
		radius:   0,
		angle:    firefly.Degrees(angle),
		angleInc: firefly.Degrees(angleInc),
		passed:   false,
	}
}

func getAngle(count int) uint32 {
	if count == 0 {
		return 135
	}
	switch count / 10 {
	case 1:
		if count%2 == 0 {
			return 180
		} else {
			return 90
		}
	case 3:
		if count%2 == 0 {
			return 270
		} else {
			return 0
		}
	case 5:
		if count%2 == 0 {
			return 225
		} else {
			return 45
		}
	default:
		return firefly.GetRandom() % 360
	}
}

// Get angle rotation speed coefficent for the given gate number.
func getAngleInc(count int) float32 {
	if count == 0 {
		return 0
	}
	switch count / 10 {
	case 0:
		return .025
	case 1:
		return 0
	case 2:
		return .033
	case 3:
		return 0
	case 4:
		return .05
	case 5:
		return 0
	case 6:
		return .075
	default:
		return .1
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
	newAngle := g.angle.Radians() + g.angleInc.Radians()
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
