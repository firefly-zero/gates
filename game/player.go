package game

import (
	"github.com/firefly-zero/firefly-go/firefly"
	"github.com/orsinium-labs/tinymath"
)

const playerR = 8

type Player struct {
	peer        firefly.Peer
	angles      [3]firefly.Angle
	anglesIndex int
}

func (p *Player) update() {
	// player := p.angles[p.anglesIndex]
	pad, touched := firefly.ReadPad(p.peer)
	if !touched {
		return
	}
	crank := pad.Azimuth()
	// for player.Sub(crank) > 180 {
	// 	player = player - 360
	// }
	// for player-crank < -180 {
	// 	player = player + 360
	// }
	// player = player + (crank.Sub(player))*0.4
	player := crank

	p.anglesIndex = p.anglesIndex + 1
	if p.anglesIndex >= len(p.angles) {
		p.anglesIndex = 0
	}
	p.angles[p.anglesIndex] = player
}

func (p *Player) render() {
	player := p.angles[p.anglesIndex]
	const distance = 60
	x := firefly.Width/2 + distance*tinymath.Cos(player.Radians())
	y := firefly.Height/2 - distance*tinymath.Sin(player.Radians())
	// p.drawTrail(player, x, y, 8)
	firefly.DrawCircle(
		firefly.P(int(x)-playerR, int(y)-playerR),
		playerR*2,
		firefly.Solid(firefly.ColorLightBlue),
	)
}
