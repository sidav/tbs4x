package main

type city struct {
	owner         *player
	name          string
	x, y          int
	maxBuildings  int
	buildingsHere []*cityBuildingStatic
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
