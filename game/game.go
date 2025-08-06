package game

import "github.com/firefly-zero/firefly-go/firefly"

var (
	font      firefly.Font
	particles *Particles
	gates     *Gates
	players   *Players
)

func Boot() {
	font = firefly.LoadFile("font", nil).Font()
	resetGame()
}

func Update() {
	particles.update()
	gates.update()
	players.update()
}

func Render() {
	firefly.ClearScreen(firefly.ColorBlack)
	particles.render()
	gates.render()
	players.render()
}

func resetGame() {
	particles = newParticles()
	gates = newGates()
	players = newPlayers()
}
