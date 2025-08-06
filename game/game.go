package game

import "github.com/firefly-zero/firefly-go/firefly"

var (
	font      firefly.Font
	particles *Particles
)

func Boot() {
	font = firefly.LoadFile("font", nil).Font()
	resetGame()
}

func Update() {
	particles.update()
}

func Render() {
	firefly.ClearScreen(firefly.ColorBlack)
	particles.render()
}

func resetGame() {
	particles = newParticles()
}
