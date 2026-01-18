package game

import "github.com/firefly-zero/firefly-go/firefly"

type Title struct {
	gates [4]*TitleGate
}

func openTitle() {
	title = &Title{gates: [...]*TitleGate{
		{firefly.Degrees(0), 1},
		{firefly.Degrees(180), 0.75},
		{firefly.Degrees(90), 0.5},
		{firefly.Degrees(270), 0.25},
	}}
}

func (t *Title) update() {
	btns := firefly.ReadButtons(firefly.Combined)
	if btns.AnyPressed() {
		resetGame()
	}

	const rotation = 2
	t.gates[0].angle = t.gates[0].angle.Add(firefly.Degrees(rotation))
	t.gates[1].angle = t.gates[1].angle.Sub(firefly.Degrees(rotation + 1))
	t.gates[2].angle = t.gates[2].angle.Add(firefly.Degrees(rotation + 2))
	t.gates[3].angle = t.gates[3].angle.Sub(firefly.Degrees(rotation + 3))
}

func (t *Title) render() {
	{
		t := "Through the Gate"
		p := firefly.P((firefly.Width-font.LineWidth(t))/2, font.CharHeight()+10)
		font.Draw(t, p, firefly.ColorWhite)
	}

	if score != nil {
		t := "score: " + formatInt(score.val)
		p := firefly.P((firefly.Width-font.LineWidth(t))/2, font.CharHeight()*2+12)
		font.Draw(t, p, firefly.ColorLightGray)
	}

	{
		t := "press E to start"
		p := firefly.P((firefly.Width-font.LineWidth(t))/2, firefly.Height-16)
		font.Draw(t, p, firefly.ColorLightGray)

		extraR := 3
		p = firefly.P(p.X+font.CharWidth()*6-extraR, p.Y-font.CharHeight()-extraR+1)
		s := firefly.Outlined(firefly.ColorLightGray, 1)
		firefly.DrawCircle(p, font.CharHeight()+extraR*2, s)
	}

	for _, g := range t.gates {
		g.render()
	}
}
