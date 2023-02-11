package main

type playerController struct {
	s                *scene
	currMode         int
	controlsPlayer   *player
	cursorX, cursorY int

	selectedUnits []*unit
}

const (
	PCMODE_NORMAL = iota
	PCMODE_UNITS_SELECTED
)

func (pc *playerController) playerControl(s *scene) {
	pc.s = s
	switch pc.currMode {
	case PCMODE_NORMAL:
		pc.normalMode()
	case PCMODE_UNITS_SELECTED:
		pc.unitsSelectedMode()
	}
	pc.normalizeCursor()
}

func (pc *playerController) normalMode() {
	key := cw.ReadKeyAsync(10)
	switch key {
	case "ESCAPE":
		GAME_RUNS = false
	case "BACKSPACE":
		pc.controlsPlayer.endedThisTurn = true
	case "TAB":
		pc.cursorX, pc.cursorY = pc.s.cities[0].x, pc.s.cities[0].y
	case "ENTER":
		unitsAtCursor := pc.s.getAllUnitsAt(pc.cursorX, pc.cursorY)
		if len(unitsAtCursor) > 0 {
			pc.selectedUnits = unitsAtCursor
			pc.currMode = PCMODE_UNITS_SELECTED
		}
	}
	vx, vy := pc.keyToDirection(key)
	pc.cursorX += vx
	pc.cursorY += vy
}

func (pc *playerController) unitsSelectedMode() {
	key := cw.ReadKeyAsync(10)
	switch key {
	case "ESCAPE", "ENTER":
		pc.currMode = PCMODE_NORMAL
	}
	vx, vy := pc.keyToDirection(key)
	if vx != 0 || vy != 0 {
		if pc.s.tryImmediateMoveUnits(pc.selectedUnits, vx, vy) {
			pc.setCursorAt(pc.selectedUnits[0].getCoords())
		}
	}
}

func (pc *playerController) setCursorAt(x, y int) {
	pc.cursorX, pc.cursorY = x, y
}

func (pc *playerController) keyToDirection(key string) (int, int) {
	switch key {
	case "UP":
		return 0, -1
	case "DOWN":
		return 0, 1
	case "LEFT":
		return -1, 0
	case "RIGHT":
		return 1, 0
	}
	return 0, 0
}

func (pc *playerController) normalizeCursor() {
	if pc.cursorX < 0 {
		pc.cursorX = 0
	}
	if pc.cursorX >= len(pc.s.tiles) {
		pc.cursorX = len(pc.s.tiles) - 1
	}
	if pc.cursorY < 0 {
		pc.cursorY = 0
	}
	if pc.cursorY >= len(pc.s.tiles[0]) {
		pc.cursorY = len(pc.s.tiles[0]) - 1
	}
}
