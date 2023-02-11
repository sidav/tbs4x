package main

import strings2 "tbs4x/lib/strings"

type unitStaticData struct {
	name            string
	figuresInUnit   int
	healthPerFigure int

	geoscapeStats unitStaticGeoscapeStats
}

type unitStaticGeoscapeStats struct {
	speed          int
	vision         int
	productionCost int
	moneyCost      int
}

func (u *unitStaticData) getCosts() (int, int) {
	return u.geoscapeStats.productionCost, u.geoscapeStats.moneyCost
}

func findUnitStaticIndexByName(name string) int {
	index := -1
	for id, u := range sTableUnits {
		if strings2.StringsAreRoughlyEqual(name, u.name) {
			if index >= 0 {
				panic("Unit search: " + name + ": many occurrences found")
			}
			index = id
		}
	}
	return index
}

var sTableUnits = []*unitStaticData{
	{
		name:            "Recon squad",
		figuresInUnit:   2,
		healthPerFigure: 2,
		geoscapeStats: unitStaticGeoscapeStats{
			speed:  2,
			vision: 2,
		},
	},
}
