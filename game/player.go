package game

import (
	"github.com/firefly-zero/firefly-go/firefly"
	"github.com/orsinium-labs/tinymath"
)

const (
	// The radius of the circle depicting the player comet.
	playerR = 8
	// The radius of the circle on which the player comet rotates.
	playerOrbit = 60
)

type Player struct {
	peer        firefly.Peer
	angles      [3]firefly.Angle
	anglesIndex int
}

func (p *Player) update() {
	player := p.angles[p.anglesIndex]
	pad, touched := firefly.ReadPad(p.peer)
	if !touched {
		return
	}
	crank := pad.Azimuth()
	for player.Sub(crank).Radians() > tinymath.Pi {
		player = player.Sub(firefly.Radians(tinymath.Tau))
	}
	for player.Sub(crank).Radians() < -tinymath.Pi {
		player = player.Add(firefly.Radians(tinymath.Tau))
	}
	delta := crank.Sub(player).Radians() * 0.4
	if !tinymath.IsNaN(delta) {
		player = player.Add(firefly.Radians(delta))
	}

	p.anglesIndex = p.anglesIndex + 1
	if p.anglesIndex >= len(p.angles) {
		p.anglesIndex = 0
	}
	p.angles[p.anglesIndex] = player
}

func (p *Player) collides(g *Gate) bool {
	player := p.angles[p.anglesIndex].Normalize().Radians()
	gate := g.angle.Normalize().Radians()
	if player < gate {
		player += tinymath.Tau
	}
	return player > gate+tinymath.Pi/2
}

func (p *Player) render() {
	player := p.angles[p.anglesIndex]
	x := firefly.Width/2 + playerOrbit*tinymath.Cos(player.Radians())
	y := firefly.Height/2 - playerOrbit*tinymath.Sin(player.Radians())
	firefly.DrawCircle(
		firefly.P(int(x)-playerR, int(y)-playerR),
		playerR*2,
		firefly.Style{
			FillColor:   firefly.ColorBlue,
			StrokeColor: firefly.ColorLightBlue,
			StrokeWidth: 1,
		},
	)
	p.drawTrail(player, x, y)
}

func (p *Player) drawTrail(player firefly.Angle, x, y float32) {
	r := player.Radians() + tinymath.Pi/2
	xLeft := x - playerR*tinymath.Cos(r)
	yLeft := y + playerR*tinymath.Sin(r)
	xRight := x + playerR*tinymath.Cos(r)
	yRight := y - playerR*tinymath.Sin(r)
	oldIndex := p.anglesIndex - (len(p.angles) - 1)
	if oldIndex < 0 {
		oldIndex = oldIndex + len(p.angles)
	}
	oldPlayer := p.angles[oldIndex]
	xTrail := x + 100*tinymath.Cos(oldPlayer.Radians())
	yTrail := y - 100*tinymath.Sin(oldPlayer.Radians())

	firefly.DrawTriangle(
		firefly.P(int(xLeft), int(yLeft)),
		firefly.P(int(xRight), int(yRight)),
		firefly.P(int(xTrail), int(yTrail)),
		firefly.Solid(firefly.ColorBlue),
	)
}
