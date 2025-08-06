package game

type Gates struct {
	delay int
	items [6]*Gate
}

func newGates() *Gates {
	return &Gates{}
}

func (ps *Gates) update() {
	for i, p := range ps.items {
		if p == nil {
			if ps.delay > 0 {
				continue
			}
			p = &Gate{}
			p.reset()
			ps.items[i] = p
			ps.delay = 45
		}
		visible := p.update()
		if !visible {
			ps.items[i] = nil
		}
	}
	ps.delay--
}

func (ps *Gates) render() {
	for _, p := range ps.items {
		if p != nil {
			p.render()
		}
	}
}
