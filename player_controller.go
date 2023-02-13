package main

import "tbs4x/lib/strings"

type playerController struct {
	s                *scene
	currMode         string
	controlsPlayer   *player
	cursorX, cursorY int

	selectedUnits []*unit
	selectedCity  *city
}

func (pc *playerController) getCursorCoords() (int, int) {
	return pc.cursorX, pc.cursorY
}

func (pc *playerController) resetMode() {
	pc.currMode = PCMODE_NORMAL
}

const (
	PCMODE_VIEW_NOTIFICATIONS = "Viewing notifications"
	PCMODE_NORMAL             = "Normal"
	PCMODE_UNITS_SELECTED     = "Unit selected"
	PCMODE_CITY_SELECTED      = "City selected"
	PCMODE_SELECTING_BLDPROD  = "Building"
	PCMODE_SELECTING_UNTPROD  = "Producing"
	PCMODE_SELECT_UNIT_ORDER  = "unit order"
)

func (pc *playerController) playerControl(s *scene) {
	pc.s = s
	switch pc.currMode {
	case PCMODE_VIEW_NOTIFICATIONS:
		pc.viewNotifications()
	case PCMODE_NORMAL:
		pc.normalMode()
	case PCMODE_UNITS_SELECTED:
		pc.unitsSelectedMode()
	case PCMODE_CITY_SELECTED:
		pc.citySelectedMode()
	case PCMODE_SELECTING_BLDPROD:
		pc.selectBuildingToMake()
	case PCMODE_SELECTING_UNTPROD:
		pc.selectUnitToMake()
	default:
		panic("No func for mode " + pc.currMode)
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
		cityAtCursor := pc.s.getCityAt(pc.getCursorCoords())
		if cityAtCursor != nil {
			pc.selectedCity = cityAtCursor
			pc.currMode = PCMODE_CITY_SELECTED
			return
		}
		unitsAtCursor := pc.s.getAllUnitsAt(pc.getCursorCoords())
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
		pc.resetMode()
	}
	vx, vy := pc.keyToDirection(key)
	if vx != 0 || vy != 0 {
		if pc.s.tryImmediateMoveUnits(pc.selectedUnits, vx, vy) {
			pc.setCursorAt(pc.selectedUnits[0].getCoords())
		}
	}
}

func (pc *playerController) viewNotifications() {
	if !pc.controlsPlayer.hasNotifications() {
		pc.resetMode()
		return
	}
	key := cw.ReadKey()
	if key == "ESCAPE" || key == "ENTER" {
		pc.resetMode()
	}
}

func (pc *playerController) citySelectedMode() {
	key := cw.ReadKey()
	if key == "ESCAPE" || key == "ENTER" {
		pc.resetMode()
	}
	if key == "b" {
		pc.currMode = PCMODE_SELECTING_BLDPROD
	}
	if key == "u" {
		pc.currMode = PCMODE_SELECTING_UNTPROD
	}
}

func (pc *playerController) selectBuildingToMake() {
	key := cw.ReadKey()
	if key == "ESCAPE" || key == "ENTER" {
		pc.currMode = PCMODE_CITY_SELECTED
	} else {
		buildables := pc.selectedCity.getListOfBuildablesHere()
		index := strings.SelectIndexFromStringsByHash(func(x int) string { return buildables[x].name }, len(buildables), key)
		if index != -1 {
			pc.selectedCity.currentProductionOrder = buildables[index]
			pc.currMode = PCMODE_CITY_SELECTED
		}
	}
}

func (pc *playerController) selectUnitToMake() {
	key := cw.ReadKey()
	if key == "ESCAPE" || key == "ENTER" {
		pc.currMode = PCMODE_CITY_SELECTED
	} else {
		buildables := pc.selectedCity.getListOfProducibleUnitsHere()
		index := strings.SelectIndexFromStringsByHash(func(x int) string { return buildables[x].name }, len(buildables), key)
		if index != -1 {
			pc.selectedCity.currentProductionOrder = buildables[index]
			pc.currMode = PCMODE_CITY_SELECTED
		}
	}
}

func (pc *playerController) selectUnitOrder() {
	key := cw.ReadKey()
	if key == "ESCAPE" || key == "ENTER" {
		pc.resetMode()
	} else {
		index := strings.SelectIndexFromStringsByHash(func(x int) string { return getNameOfOrder(x) }, ORDERS_COUNT, key)
		if index != -1 {
			pc.resetMode()
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
