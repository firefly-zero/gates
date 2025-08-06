package game

import "github.com/firefly-zero/firefly-go/firefly"

func randomAngle() firefly.Angle {
	return firefly.Degrees(float32(firefly.GetRandom() % 360))
}

func randomUint(min, max uint32) uint32 {
	return firefly.GetRandom()%(max-min) + min
}

func randomInt(min, max int) int {
	return int(firefly.GetRandom()%uint32(max-min)) + min
}
