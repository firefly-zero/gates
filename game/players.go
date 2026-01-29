package game

import "github.com/firefly-zero/firefly-go/firefly"

type Players struct {
	items []*Player
}

func newPlayers() *Players {
	peers := firefly.GetPeers().Slice()
	me := firefly.GetMe()

	// Put the current player last on the list
	// so that it is always rendered on top.
	for i, peer := range peers {
		if peer == me {
			lastIdx := len(peers) - 1
			last := peers[lastIdx]
			peers[lastIdx] = peer
			peers[i] = last
		}
	}

	players := make([]*Player, len(peers))
	for i, peer := range peers {
		players[i] = &Player{
			me:   peer == me,
			peer: peer,
		}
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
			openTitle()
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
