package game

import (
	"github.com/firefly-zero/firefly-go/firefly"
	"github.com/orsinium-labs/tinymath"
)

type Particle struct {
	angle  firefly.Angle
	pos    float32
	posInc float32
	scale  float32
}

func (p *Particle) reset() {
	p.angle = randomAngle()
	p.pos = 20
	p.posInc = 2
	p.scale = float32(randomUint(25, 100)) * 0.01
}

func (p *Particle) update() bool {
	p.pos = p.pos + p.posInc
	p.posInc = p.posInc * 1.1
	return p.pos <= firefly.Width/2
}

func (p *Particle) render() {
	x := firefly.Width/2 + p.pos*tinymath.Cos(p.angle.Radians())
	y := firefly.Height/2 + p.pos*tinymath.Sin(p.angle.Radians())
	distance := p.posInc * (p.scale * 2)
	xNext := x + distance*tinymath.Cos(p.angle.Radians())
	yNext := y + distance*tinymath.Sin(p.angle.Radians())

	line := firefly.LineStyle{
		Color: firefly.ColorDarkBlue,
		Width: int(distance * (p.scale * 0.5)),
	}
	line.Draw(
		firefly.P(int(x), int(y)),
		firefly.P(int(xNext), int(yNext)),
	)

}
