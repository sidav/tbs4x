package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
)

// View not separated from model... bad practice?
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
		fmt.Sprintf("Space %d/%d", 0, city.maxBuildings),
		fmt.Sprintf("Buildings here: NONE"),
		fmt.Sprintf("Current production: NONE (p to change)"),
	}
	cw.ResetStyle()
	for i, l := range lines {
		cw.PutString(l, x+1, y+i+1)
	}
}
