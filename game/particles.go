package game

type Particles struct {
	delay int
	items [12]*Particle
}

func newParticles() *Particles {
	return &Particles{}
}

func (ps *Particles) update() {
	for i, p := range ps.items {
		if p == nil {
			if ps.delay != 0 {
				continue
			}
			p = &Particle{}
			p.reset()
			ps.items[i] = p
			ps.delay = 2
		}
		visible := p.update()
		if !visible {
			p.reset()
		}
	}
	ps.delay--
}

func (ps *Particles) render() {
	for _, p := range ps.items {
		if p != nil {
			p.render()
		}
	}
}
