package game

import (
	"unsafe"

	"github.com/firefly-zero/firefly-go/firefly"
)

type Score struct {
	val int
}

func newScore() *Score {
	return &Score{}
}

func (s *Score) inc() {
	s.val += 1
}

func (s *Score) render() {
	t := formatInt(s.val)
	p := firefly.P(
		(firefly.Width-font.LineWidth(t))/2,
		font.CharHeight()+2,
	)
	font.Draw(t, p, firefly.ColorLightGray)
}

func formatInt(i int) string {
	buf := []byte{'0' + byte(i/10), '0' + byte(i%10)}
	return unsafe.String(&buf[0], 2)
}
