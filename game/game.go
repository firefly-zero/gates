package game

import "github.com/firefly-zero/firefly-go/firefly"

var (
	font      firefly.Font
	particles *Particles
	gates     *Gates
	players   *Players
	score     *Score
	title     *Title
)

func Boot() {
	font = firefly.LoadFile("font", nil).Font()
	openTitle()
}

func Update() {
	if title != nil {
		title.update()
		return
	}
	particles.update()
	gates.update()
	players.update()
}

func Render() {
	firefly.ClearScreen(firefly.ColorBlack)
	if title != nil {
		title.render()
		return
	}
	particles.render()
	gates.render()
	score.render()
	players.render()
}

func resetGame() {
	particles = newParticles()
	gates = newGates()
	players = newPlayers()
	score = newScore()
	title = nil
}
