package state

import "testing"

func TestState(t *testing.T) {
	start := StartState{}
	game := GameContext{
		Next: &start,
	}

	i := 0
	for game.Next.executeState(&game) {
		i++
		if i == 10000 {
			break
		}
	}
}
