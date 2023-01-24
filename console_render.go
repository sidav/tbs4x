package main

import "github.com/gdamore/tcell/v2"

type asciiRenderer struct {
	camX, camY   int
	tileW, tileH int
	pc           *playerController
}

func newAsciiRenderer() *asciiRenderer {
	return &asciiRenderer{
		camX:  0,
		camY:  0,
		tileW: 4,
		tileH: 3,
	}
}

func (rs *asciiRenderer) renderMainScreen(s *scene, pc *playerController) {
	rs.pc = pc
	cw.ClearScreen()
	consW, consH := cw.GetConsoleSize()
	rs.camX = pc.cursorX - (consW/rs.tileW)/2
	rs.camY = pc.cursorY - (consH/rs.tileH)/2
	for x := range s.tiles {
		for y := range s.tiles[x] {
			sx, sy := rs.globalToOnScreen(x, y)
			rs.renderTile(s.tiles[x][y], sx, sy)
		}
	}
	rs.renderCursor()
	cw.FlushScreen()
}

func (rs *asciiRenderer) renderTile(t *tile, sx, sy int) {
	char := '?'
	switch t.code {
	case TILE_WATER:
		cw.SetStyle(tcell.ColorBlue, tcell.ColorBlack)
		char = '~'
	case TILE_ROCK:
		cw.SetStyle(tcell.ColorDarkRed, tcell.ColorBlack)
		char = '^'
	case TILE_SAND:
		cw.SetStyle(tcell.ColorYellow, tcell.ColorBlack)
		char = '.'
	}
	for x := sx; x < sx+rs.tileW; x++ {
		for y := sy; y < sy+rs.tileH; y++ {
			cw.PutChar(char, x, y)
		}
	}
}

func (rs *asciiRenderer) renderCursor() {
	sx, sy := rs.globalToOnScreen(rs.pc.cursorX, rs.pc.cursorY)
	cw.SetStyle(tcell.ColorDarkRed, tcell.ColorBlack)
	cw.PutChar('|', sx-1, sy)
	cw.PutChar('|', sx+rs.tileW, sy)
	cw.PutChar('|', sx-1, sy+rs.tileH-1)
	cw.PutChar('|', sx+rs.tileW, sy+rs.tileH-1)

	cw.PutChar('-', sx, sy-1)
	cw.PutChar('-', sx+rs.tileW-1, sy-1)
	cw.PutChar('-', sx, sy+rs.tileH)
	cw.PutChar('-', sx+rs.tileW-1, sy+rs.tileH)
}

func (rs *asciiRenderer) globalToOnScreen(gx, gy int) (int, int) {
	return rs.tileW * (gx - rs.camX), rs.tileH * (gy - rs.camY)
}

func (rs *asciiRenderer) onScreenToGlobal(sx, sy int) (int, int) {
	return rs.camX + sx, rs.camY + sy
}
