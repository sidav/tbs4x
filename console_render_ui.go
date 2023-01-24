package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
)

func (rs *asciiRenderer) renderUI() {
	rs.renderCursor()
	rs.drawCenteredStringAndIncrementLine("Turn 1", rs.uiPanelCenterX)
	rs.drawCenteredStringAndIncrementLine(fmt.Sprintf("(%d, %d)", rs.pc.cursorX, rs.pc.cursorY), rs.uiPanelCenterX)
	rs.currUiLine++
	rs.renderUnderCursorInfo()
}

func (rs *asciiRenderer) renderUnderCursorInfo() {
	cityHere := rs.sc.getCityAt(rs.pc.cursorX, rs.pc.cursorY)
	if cityHere != nil {
		cw.SetFg(tcell.ColorWhite)
		rs.drawCenteredStringAndIncrementLine(cityHere.name, rs.uiPanelCenterX)
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

func (rs *asciiRenderer) drawStringAndIncrementLine(str string, x int) {
	cw.PutString(str, x, rs.currUiLine)
	rs.currUiLine++
}

func (rs *asciiRenderer) drawCenteredStringAndIncrementLine(str string, x int) {
	cw.PutStringCenteredAt(str, x, rs.currUiLine)
	rs.currUiLine++
}
