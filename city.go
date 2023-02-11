package main

type city struct {
	owner         *player
	name          string
	x, y          int
	maxBuildings  int
	buildingsHere []*cityBuildingStatic

	currentProductionOrder producible
	accumulatedProduction  int
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

func (c *city) getTotalProductionPowerForType(ptype int) int {
	prod := 0
	for _, b := range c.buildingsHere {
		for _, pr := range b.prodPowers {
			if pr.prodCode == ptype {
				prod += pr.prodPower
			}
		}
	}
	return prod
}

func (c *city) getETAForProducing(p producible) int {
	prodForType := c.getTotalProductionPowerForType(p.getProductionTypeRequired())
	if prodForType == 0 {
		return 999
	}
	eta := (p.getProductionCost() - c.accumulatedProduction) / prodForType
	if eta == 0 || (p.getProductionCost()-c.accumulatedProduction)%prodForType > 0 {
		eta++
	}
	return eta
}

func (c *city) getETAForCurrentProduction() int {
	if c.currentProductionOrder == nil {
		return 0
	}
	return c.getETAForProducing(c.currentProductionOrder)
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

func (c *city) getListOfProducibleUnitsHere() []*unitStaticData {
	arr := make([]*unitStaticData, 0)
	for _, u := range sTableUnits {
		if c.getTotalProductionPowerForType(u.geoscapeStats.producedWithType) > 0 {
			arr = append(arr, u)
		}
	}
	return arr
}
