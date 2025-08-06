package game

type Gates struct {
	// Time in frames until the next gate is generated.
	delay int
	// The total number of gates generated during the game.
	count int
	items [6]*Gate
	// The gate currently crossing the player's orbit.
	current *Gate
}

func newGates() *Gates {
	return &Gates{}
}

func (gs *Gates) update() {
	gs.current = nil
	for i, g := range gs.items {
		if g == nil {
			if gs.delay > 0 {
				continue
			}
			g = newGate(gs.count)
			gs.items[i] = g
			gs.delay = 45
			gs.count++
		}
		visible := g.update()
		if !visible {
			gs.items[i] = nil
		}
		if !g.passed && g.radius >= playerOrbit {
			gs.current = g
		}
	}
	gs.delay--
}

func (gs *Gates) render() {
	for _, g := range gs.items {
		if g != nil {
			g.render()
		}
	}
}
