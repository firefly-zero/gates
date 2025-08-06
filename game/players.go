package game

import "github.com/firefly-zero/firefly-go/firefly"

type Players struct {
	items []*Player
}

func newPlayers() *Players {
	peers := firefly.GetPeers()
	players := make([]*Player, peers.Len())
	for i, peer := range peers.Slice() {
		players[i] = &Player{peer: peer}
	}
	return &Players{items: players}
}

func (ps *Players) update() {
	for _, p := range ps.items {
		p.update()
	}
	g := gates.current
	if g != nil {
		if ps.collides(g) {
			// TODO: game over
			score.val = 0
		} else {
			g.passed = true
			score.inc()
		}
	}
}

func (ps *Players) collides(g *Gate) bool {
	for _, p := range ps.items {
		if p != nil {
			if p.collides(g) {
				return true
			}
		}
	}
	return false
}

func (ps *Players) render() {
	for _, p := range ps.items {
		if p != nil {
			p.render()
		}
	}
}
