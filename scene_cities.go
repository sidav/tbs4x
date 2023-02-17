package main

import "fmt"

func (s *scene) performProductionForCity(c *city) {
	if c.currentProductionOrder == nil {
		return
	}
	c.accumulatedProduction += c.getTotalProductionPowerForType(c.currentProductionOrder.getProductionTypeRequired())
	prodCost := c.currentProductionOrder.getProductionCost()
	if c.accumulatedProduction >= prodCost {
		if b, ok := c.currentProductionOrder.(*cityBuildingStatic); ok {
			c.addBuilding(b)
		}
		if u, ok := c.currentProductionOrder.(*unitStaticData); ok {
			s.addUnit(findUnitStaticIndexByName(u.name), c.owner, c.x, c.y+1)
		}
		c.owner.addNotification(fmt.Sprintf("%s finished %s %s",
			c.name,
			getProductionTypeString(c.currentProductionOrder.getProductionTypeRequired()),
			c.currentProductionOrder.getName()))
		c.currentProductionOrder = nil
		c.accumulatedProduction = 0
	}
}
