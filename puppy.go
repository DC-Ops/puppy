package puppy

import (
	"github.com/DC-Ops/gshep"
)

func BasicBark() string {
	sound := "Woof! Woof!"
	return sound
}

func GshepBark() string {
	sound := gshep.GshepBark()
	return sound
}
