package game

import "github.com/firefly-zero/firefly-go/firefly"

type Title struct{}

func openTitle() {
	title = &Title{}
}

func (Title) update() {
	btns := firefly.ReadButtons(firefly.Combined)
	if btns.S {
		resetGame()
	}
}

func (Title) render() {
	{
		t := "Through the Gate"
		p := firefly.P((firefly.Width-font.LineWidth(t))/2, font.CharHeight()+10)
		font.Draw(t, p, firefly.ColorWhite)
	}

	{
		t := "score: " + formatInt(score.val)
		p := firefly.P((firefly.Width-font.LineWidth(t))/2, font.CharHeight()*2+12)
		font.Draw(t, p, firefly.ColorLightGray)
	}

	{
		t := "press S to start"
		p := firefly.P((firefly.Width-font.LineWidth(t))/2, firefly.Height-4)
		font.Draw(t, p, firefly.ColorLightGray)
	}
}
