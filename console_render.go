package main

import "github.com/gdamore/tcell/v2"

type asciiRenderer struct {
	consW, consH int
	tileW, tileH int

	camX, camY int

	uiPanelW       int
	uiPanelCenterX int
	currUiLine     int

	pc *playerController
	sc *scene
}

func newAsciiRenderer() *asciiRenderer {
	return &asciiRenderer{
		camX:  0,
		camY:  0,
		tileW: 4,
		tileH: 3,
	}
}

func (rs *asciiRenderer) updateVars(s *scene, pc *playerController) {
	rs.sc = s
	rs.pc = pc
	rs.currUiLine = 0
	cw.ClearScreen()
	rs.consW, rs.consH = cw.GetConsoleSize()
	rs.uiPanelCenterX = rs.consW - rs.uiPanelW/2
	rs.uiPanelW = rs.consW / 4
	rs.camX = pc.cursorX - ((rs.consW-rs.uiPanelW)/rs.tileW)/2
	rs.camY = pc.cursorY - (rs.consH/rs.tileH)/2
}

func (rs *asciiRenderer) renderMainScreen(s *scene, pc *playerController) {
	rs.updateVars(s, pc)

	for x := range s.tiles {
		for y := range s.tiles[x] {
			if rs.areGlobalCoordsOnScreen(x, y) {
				sx, sy := rs.globalToOnScreen(x, y)
				rs.renderTile(s.tiles[x][y], sx, sy)
			}
		}
	}
	for _, c := range s.cities {
		rs.renderCity(c)
	}
	rs.renderUI()
	cw.FlushScreen()
}

func (rs *asciiRenderer) renderTile(t *tile, sx, sy int) {
	char := '?'
	switch t.code {
	case TILE_WATER:
		cw.SetStyle(tcell.ColorBlue, tcell.ColorBlack)
		char = '~'
	case TILE_MOUNTAIN:
		cw.SetStyle(tcell.ColorDarkRed, tcell.ColorBlack)
		char = '^'
	case TILE_GRASS:
		cw.SetStyle(tcell.ColorDarkGreen, tcell.ColorBlack)
		char = '.'
	case TILE_SAND:
		cw.SetStyle(tcell.ColorYellow, tcell.ColorBlack)
		char = '.'
	}
	for x := sx; x < sx+rs.tileW; x++ {
		for y := sy; y < sy+rs.tileH; y++ {
			cw.PutChar(char, x, y)
		}
	}
	if t.resourceAmountHere > 0 {
		switch t.resourceCode {
		case RES_GOLD:
			cw.SetFg(tcell.ColorYellow)
		case RES_GREENIUM:
			cw.SetFg(tcell.ColorGreen)
		}
		for i := 0; i < rs.tileW*rs.tileH; i += 3 {
			cw.PutChar('*', sx+((i+t.resourceAmountHere)%rs.tileW), sy+(i/rs.tileW))
		}
	}
}

func (rs *asciiRenderer) renderCity(c *city) {
	if !rs.areGlobalCoordsOnScreen(c.x, c.y) {
		return
	}
	sx, sy := rs.globalToOnScreen(c.x, c.y)
	cityImage := []string{
		"/=^\\",
		"=&|=",
		"\\==/",
	}
	cw.SetFg(tcell.ColorWhite)
	for x := 0; x < rs.tileW; x++ {
		for y := 0; y < rs.tileH; y++ {
			cw.PutChar(rune(cityImage[y][x]), sx+x, sy+y)
		}
	}
	cw.PutStringCenteredAt(c.name, sx+rs.tileW/2, sy+rs.tileH)
}

func (rs *asciiRenderer) areGlobalCoordsOnScreen(gx, gy int) bool {
	return gx-rs.camX < (rs.consW-rs.uiPanelW)/rs.tileW &&
		gy-rs.camY <= rs.consH/rs.tileH
}

func (rs *asciiRenderer) globalToOnScreen(gx, gy int) (int, int) {
	return rs.tileW * (gx - rs.camX), rs.tileH * (gy - rs.camY)
}

func (rs *asciiRenderer) onScreenToGlobal(sx, sy int) (int, int) {
	return rs.camX + sx, rs.camY + sy
}
