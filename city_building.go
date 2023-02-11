package main

import (
	strings2 "tbs4x/lib/strings"
)

type cityBuildingStatic struct {
	name string

	unbuildable      bool
	productionCost   int
	producedWithType int
	moneyCost        int

	size                    int
	requiresBuildingsInCity []string

	prodPowers []productionAbility
}

func (cb *cityBuildingStatic) getCosts() (int, int) {
	return cb.productionCost, cb.moneyCost
}

func (cb *cityBuildingStatic) getProductionTypeRequired() int {
	return cb.producedWithType
}

func (cb *cityBuildingStatic) getName() string {
	return cb.name
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
		name:        "Frontal HQ",
		unbuildable: true,
		size:        5,
		prodPowers: []productionAbility{{
			prodCode:  PRODUCTION_BUILDING,
			prodPower: 5,
		}},
	},
	{
		name:           "Construction Yard",
		productionCost: 0,
		moneyCost:      0,
		size:           4,
		prodPowers: []productionAbility{{
			prodCode:  PRODUCTION_BUILDING,
			prodPower: 2,
		}},
	},
	{
		name:             "Barracks",
		productionCost:   25,
		moneyCost:        20,
		size:             2,
		producedWithType: PRODUCTION_BUILDING,
	},
	{
		name:             "Refinery",
		productionCost:   50,
		moneyCost:        30,
		size:             3,
		producedWithType: PRODUCTION_BUILDING,
	},
}
