package main

type city struct {
	owner         *player
	name          string
	x, y          int
	maxBuildings  int
	buildingsHere []*cityBuildingStatic

	currentProductionOrder       producible
	currentAccumulatedProduction int
}

func (c *city) addBuilding(s *cityBuildingStatic) {
	c.buildingsHere = append(c.buildingsHere, s)
}

func (c *city) countUsedSpace() int {
	space := 0
	for _, b := range c.buildingsHere {
		space += b.size
	}
	return space
}

func (c *city) getListOfBuildablesHere() []*cityBuildingStatic {
	arr := make([]*cityBuildingStatic, 0)
	for _, b := range sTableBuildings {
		if b.unbuildable {
			continue
		}
		arr = append(arr, b)
	}
	return arr
}
