package main

import strings2 "tbs4x/lib/strings"

type unitStaticData struct {
	name            string
	figuresInUnit   int
	healthPerFigure int

	geoscapeStats unitStaticGeoscapeStats
}

type unitStaticGeoscapeStats struct {
	speed            int
	vision           int
	productionCost   int
	producedWithType int
	moneyCost        int
}

func (ud *unitStaticData) getProductionCost() int {
	return ud.geoscapeStats.productionCost
}

func (ud *unitStaticData) getMoneyCost() int {
	return ud.geoscapeStats.moneyCost
}

func (ud *unitStaticData) getProductionTypeRequired() int {
	return ud.geoscapeStats.producedWithType
}

func (ud *unitStaticData) getName() string {
	return ud.name
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

			productionCost:   10,
			producedWithType: PRODUCTION_GROUND_MECH,
			moneyCost:        10,
		},
	},
	{
		name:            "GI Squad",
		figuresInUnit:   5,
		healthPerFigure: 2,
		geoscapeStats: unitStaticGeoscapeStats{
			speed:  1,
			vision: 1,

			productionCost:   10,
			producedWithType: PRODUCTION_INFANTRY,
			moneyCost:        10,
		},
	},
}
