package main

func (s *scene) performProductionForCity(c *city) {
	if c.currentProductionOrder == nil {
		return
	}
	c.currentAccumulatedProduction += c.getTotalProductionPowerForType(c.currentProductionOrder.getProductionTypeRequired())
	prodCost := c.currentProductionOrder.getProductionCost()
	if c.currentAccumulatedProduction >= prodCost {
		if b, ok := c.currentProductionOrder.(*cityBuildingStatic); ok {
			c.addBuilding(b)
		}
		if u, ok := c.currentProductionOrder.(*unitStaticData); ok {
			newUnt := &unit{
				owner: c.owner,
				id:    findUnitStaticIndexByName(u.name),
				x:     c.x,
				y:     c.y + 1,
			}
			newUnt.initByStatic()
			s.addUnit(newUnt)
		}
		c.currentProductionOrder = nil
		c.currentAccumulatedProduction = 0
	}
}
