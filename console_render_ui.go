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
	if rs.sc.tiles[rs.pc.cursorX][rs.pc.cursorY].resourceAmountHere > 0 {
		cw.SetFg(tcell.ColorYellow)
		rs.drawCenteredStringAndIncrementLine(fmt.Sprintf("Resources: %d", rs.sc.tiles[rs.pc.cursorX][rs.pc.cursorY].resourceAmountHere), rs.uiPanelCenterX)
	}

	cityHere := rs.sc.getCityAt(rs.pc.cursorX, rs.pc.cursorY)
	if cityHere != nil {
		cw.SetFg(tcell.ColorWhite)
		rs.drawCenteredStringAndIncrementLine(cityHere.name, rs.uiPanelCenterX)
		rs.drawCenteredStringAndIncrementLine(fmt.Sprintf(" Size: %d", cityHere.maxBuildings), rs.uiPanelCenterX)
	}
	rs.currUiLine++

	unitsHere := rs.sc.getAllUnitsAt(rs.pc.cursorX, rs.pc.cursorY)
	if len(unitsHere) > 0 {
		rs.drawCenteredStringAndIncrementLine(fmt.Sprintf("Squad: %d MP", arrayOfUnits(unitsHere).getMinMovementPoints()),
			rs.uiPanelCenterX)
	}
	for i, u := range unitsHere {
		cw.SetFg(tcell.ColorWhite)
		rs.drawCenteredStringAndIncrementLine(fmt.Sprintf("%d. %s", i+1, u.getStaticData().name),
			rs.uiPanelCenterX)

	}
}

func (rs *asciiRenderer) renderCursor() {
	sx, sy := rs.globalToOnScreen(rs.pc.cursorX, rs.pc.cursorY)
	switch rs.pc.currMode {
	case PCMODE_NORMAL:
		cw.SetStyle(tcell.ColorDarkRed, tcell.ColorBlack)
		cw.PutChar('|', sx-1, sy)
		cw.PutChar('|', sx+rs.tileW, sy)
		cw.PutChar('|', sx-1, sy+rs.tileH-1)
		cw.PutChar('|', sx+rs.tileW, sy+rs.tileH-1)

		cw.PutChar('-', sx, sy-1)
		cw.PutChar('-', sx+rs.tileW-1, sy-1)
		cw.PutChar('-', sx, sy+rs.tileH)
		cw.PutChar('-', sx+rs.tileW-1, sy+rs.tileH)
	case PCMODE_UNITS_SELECTED:
		cw.SetStyle(tcell.ColorGreen, tcell.ColorBlack)
		cw.PutChar('<', sx-1, sy+1)
		cw.PutChar('<', sx-1, sy+rs.tileH-2)
		cw.PutChar('>', sx+rs.tileW, sy+1)
		cw.PutChar('>', sx+rs.tileW, sy+rs.tileH-2)

		cw.PutChar('^', sx+1, sy-1)
		cw.PutChar('^', sx+rs.tileW-2, sy-1)
		cw.PutChar('v', sx+1, sy+rs.tileH)
		cw.PutChar('v', sx+rs.tileW-2, sy+rs.tileH)
	}
}

func (rs *asciiRenderer) drawStringAndIncrementLine(str string, x int) {
	cw.PutString(str, x, rs.currUiLine)
	rs.currUiLine++
}

func (rs *asciiRenderer) drawCenteredStringAndIncrementLine(str string, x int) {
	cw.PutStringCenteredAt(str, x, rs.currUiLine)
	rs.currUiLine++
}
