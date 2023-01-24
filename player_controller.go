package main

type playerController struct {
	cursorX, cursorY int
}

func (pc *playerController) playerControl() {
	key := cw.ReadKeyAsync(10)
	switch key {
	case "ESCAPE":
		GAME_RUNS = false
	case "UP":
		pc.cursorY--
	case "DOWN":
		pc.cursorY++
	case "LEFT":
		pc.cursorX--
	case "RIGHT":
		pc.cursorX++
	}
}
