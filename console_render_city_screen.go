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
		prodLine = getProductionTypeString(city.currentProductionOrder.getProductionTypeRequired()) + " " + city.currentProductionOrder.getName()
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
	cw.DrawFilledRect(' ', x, y, w, h)
	cw.SetStyle(tcell.ColorWhite, tcell.ColorBlue)
	cw.DrawRect(x, y, w, h)
	cw.SetStyle(tcell.ColorBlue, tcell.ColorBlack)
	cw.PutStringCenteredAt(" Select building ", x+w/2, y)

	lines := []string{
		fmt.Sprintf("City at %d, %d", city.x, city.y),
		fmt.Sprintf("Buildings here (space %d/%d)", city.countUsedSpace(), city.maxBuildings),
		"",
	}
	buildable := city.getListOfBuildablesHere()
	hashes := strings.HashStringsToShortestDistincts(func(x int) string { return buildable[x].name }, len(buildable))
	for i, b := range buildable {
		lines = append(lines,
			fmt.Sprintf("%s - %s", hashes[i], b.name),
			fmt.Sprintf(" $%d, size %d", b.moneyCost, b.size),
		)
	}

	cw.ResetStyle()
	for i, l := range lines {
		cw.PutString(l, x+1, y+i+1)
	}
}

func (r *asciiRenderer) showAvailableUnitsToMake() {
	city := r.pc.selectedCity
	const offsetW = 8
	const offsetH = 12
	x, y := r.consW/offsetW, r.consH/offsetH
	w, h := r.consW-(2*r.consW/offsetW), r.consH-(2*r.consH/offsetH)
	cw.ResetStyle()
	cw.DrawFilledRect(' ', x, y, w, h)
	cw.SetStyle(tcell.ColorWhite, tcell.ColorBlue)
	cw.DrawRect(x, y, w, h)
	cw.SetStyle(tcell.ColorBlue, tcell.ColorBlack)
	cw.PutStringCenteredAt(" Select building ", x+w/2, y)

	lines := []string{
		fmt.Sprintf("City at %d, %d", city.x, city.y),
	}
	buildable := city.getListOfProducibleUnitsHere()
	hashes := strings.HashStringsToShortestDistincts(func(x int) string { return buildable[x].name }, len(buildable))
	for i, b := range buildable {
		lines = append(lines,
			fmt.Sprintf("%s - %s", hashes[i], b.name),
			fmt.Sprintf(" $%d", b.getMoneyCost()),
		)
	}

	cw.ResetStyle()
	for i, l := range lines {
		cw.PutString(l, x+1, y+i+1)
	}
}
