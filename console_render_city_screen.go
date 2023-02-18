package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"tbs4x/lib/strings"
)

func (r *asciiRenderer) showCityScreen() {
	city := r.pc.selectedCity
	const offsetPart = 8
	x, y := r.consW/offsetPart, r.consH/offsetPart
	w, h := r.consW-(2*r.consW/offsetPart), r.consH-(2*r.consH/offsetPart)
	cw.ResetStyle()
	cw.DrawFilledRect(' ', x, y, w, h)
	cw.SetStyle(tcell.ColorWhite, tcell.ColorBlue)
	cw.DrawRect(x, y, w, h)
	cw.SetStyle(tcell.ColorBlue, tcell.ColorBlack)
	cw.PutStringCenteredAt(" "+city.name+" ", x+w/2, y)

	lines := []string{
		fmt.Sprintf("City at %d, %d", city.x, city.y),
		fmt.Sprintf("Buildings here (space %d/%d)", city.countUsedSpace(), city.maxBuildings),
	}
	for _, b := range city.buildingsHere {
		lines = append(lines, fmt.Sprintf("  %s (size %d)", b.name, b.size))
	}
	prodLine := "none"
	if city.currentProductionOrder != nil {
		prodLine = fmt.Sprintf("%s %s (ETA %d turns)",
			getProductionTypeString(city.currentProductionOrder.getProductionTypeRequired()),
			city.currentProductionOrder.getName(), city.getETAForCurrentProduction())
	}
	lines = append(lines,
		"",
		fmt.Sprintf("Current production: %s", prodLine),
		fmt.Sprintf("b to change to building, u to change to unit"),
	)

	cw.ResetStyle()
	for i, l := range lines {
		cw.PutString(l, x+1, y+i+1)
	}
}

func (r *asciiRenderer) showAvailableBuildingsToMake() {
	city := r.pc.selectedCity
	const offsetW = 8
	const offsetH = 12
	x, y := r.consW/offsetW, r.consH/offsetH
	w, h := r.consW-(2*r.consW/offsetW), r.consH-(2*r.consH/offsetH)
	cw.ResetStyle()
	r.drawMenuRect("Select building", x, y, w, h)

	r.currUiLine = y + 1
	cw.SetFg(tcell.ColorWhite)
	r.drawStringAndIncrementLine(fmt.Sprintf("City at %d, %d", city.x, city.y), x+1)
	r.drawStringAndIncrementLine(fmt.Sprintf("Buildings here (space %d/%d)", city.countUsedSpace(), city.maxBuildings), x+1)
	r.drawStringAndIncrementLine("", x)

	buildable := city.getListOfBuildablesHere()
	hashes := strings.HashStringsToShortestDistincts(func(x int) string { return buildable[x].name }, len(buildable))
	for i, b := range buildable {
		cw.SetFg(tcell.ColorWhite)
		if r.pc.controlsPlayer.currentMoney < b.moneyCost {
			cw.SetFg(tcell.ColorDarkGray)
		}
		r.drawStringAndIncrementLine(fmt.Sprintf("%s - %s", hashes[i], b.name), x+1)
		r.drawStringAndIncrementLine(fmt.Sprintf(" $%d, size %d, ETA %d turns", b.moneyCost, b.size, city.getETAForProducing(b)), x+1)
	}
}

func (r *asciiRenderer) showAvailableUnitsToMake() {
	city := r.pc.selectedCity
	const offsetW = 8
	const offsetH = 12
	x, y := r.consW/offsetW, r.consH/offsetH
	w, h := r.consW-(2*r.consW/offsetW), r.consH-(2*r.consH/offsetH)
	cw.ResetStyle()
	r.drawMenuRect("Select unit", x, y, w, h)

	buildable := city.getListOfProducibleUnitsHere()
	hashes := strings.HashStringsToShortestDistincts(func(x int) string { return buildable[x].name }, len(buildable))

	r.currUiLine = y + 1
	cw.SetFg(tcell.ColorWhite)
	r.drawStringAndIncrementLine(fmt.Sprintf("City at %d, %d", city.x, city.y), x+1)

	for i, b := range buildable {
		cw.SetFg(tcell.ColorWhite)
		if r.pc.controlsPlayer.currentMoney < b.geoscapeStats.moneyCost {
			cw.SetFg(tcell.ColorDarkGray)
		}
		r.drawStringAndIncrementLine(fmt.Sprintf("%s - %s", hashes[i], b.name), x+1)
		r.drawStringAndIncrementLine(fmt.Sprintf(" $%d ETA %d turns", b.geoscapeStats.moneyCost, city.getETAForProducing(b)), x+1)
	}
}

func (r *asciiRenderer) drawMenuRect(title string, x, y, w, h int) {
	cw.DrawFilledRect(' ', x, y, w, h)
	cw.SetStyle(tcell.ColorWhite, tcell.ColorBlue)
	cw.DrawRect(x, y, w, h)
	cw.SetStyle(tcell.ColorBlue, tcell.ColorBlack)
	cw.PutStringCenteredAt(" "+title+" ", x+w/2, y)
}
