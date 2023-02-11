package main

import (
	strings2 "tbs4x/lib/strings"
)

type cityBuildingStatic struct {
	name string

	productionCost int
	moneyCost      int

	size                    int
	requiresBuildingsInCity []string

	buildPower int
}

func (cb *cityBuildingStatic) getCosts() (int, int) {
	return cb.productionCost, cb.moneyCost
}

func findBuildingInTableByName(name string) *cityBuildingStatic {
	var foundBld *cityBuildingStatic
	for _, b := range sTableBuildings {
		if strings2.StringsAreRoughlyEqual(name, b.name) {
			if foundBld != nil {
				panic("Building search: " + name + ": many occurrences found")
			}
			foundBld = b
		}
	}
	return foundBld
}

var sTableBuildings = []*cityBuildingStatic{
	{
		name:           "Frontal HQ",
		productionCost: 0,
		moneyCost:      0,
		size:           5,
	},
	{
		name:           "Construction Yard",
		productionCost: 0,
		moneyCost:      0,
		size:           4,
	},
	{
		name:           "Barracks",
		productionCost: 25,
		moneyCost:      20,
		size:           2,
	},
	{
		name:           "Refinery",
		productionCost: 50,
		moneyCost:      30,
		size:           3,
	},
}
