package game

import (
	"math"

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
	for player.Sub(crank).Radians() > math.Pi {
		player = player.Sub(firefly.Radians(math.Pi * 2))
	}
	for player.Sub(crank).Radians() < -math.Pi {
		player = player.Add(firefly.Radians(math.Pi * 2))
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
	player := p.angles[p.anglesIndex].Add(firefly.Radians(math.Pi))
	for player.Add(g.angle).Radians() > math.Pi {
		player = player.Sub(firefly.Radians(math.Pi * 2))
	}
	for player.Sub(g.angle).Radians() < -math.Pi {
		player = player.Add(firefly.Radians(math.Pi * 2))
	}
	if player.Radians() < g.angle.Radians() {
		return true
	}
	if player.Radians() > g.angle.Radians()+math.Pi {
		return true
	}
	return false
}

func (p *Player) render() {
	player := p.angles[p.anglesIndex]
	x := firefly.Width/2 + playerOrbit*tinymath.Cos(player.Radians())
	y := firefly.Height/2 - playerOrbit*tinymath.Sin(player.Radians())
	p.drawTrail(player, x, y)
	firefly.DrawCircle(
		firefly.P(int(x)-playerR, int(y)-playerR),
		playerR*2,
		firefly.Style{
			FillColor:   firefly.ColorBlue,
			StrokeColor: firefly.ColorLightBlue,
			StrokeWidth: 1,
		},
	)
}

func (p *Player) drawTrail(player firefly.Angle, x, y float32) {
	r := player.Radians() + math.Pi/2
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
