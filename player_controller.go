package main

type playerController struct {
	controlsPlayer   *player
	cursorX, cursorY int
}

func (pc *playerController) playerControl(s *scene) {
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
	case "TAB":
		pc.cursorX, pc.cursorY = s.cities[0].x, s.cities[0].y
	}
	if pc.cursorX < 0 {
		pc.cursorX = 0
	}
	if pc.cursorX >= len(s.tiles) {
		pc.cursorX = len(s.tiles) - 1
	}
	if pc.cursorY < 0 {
		pc.cursorY = 0
	}
	if pc.cursorY >= len(s.tiles[0]) {
		pc.cursorY = len(s.tiles[0]) - 1
	}
}
