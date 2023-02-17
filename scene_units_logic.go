package main

import (
	"tbs4x/lib/calculations"
)

func (s *scene) getUnitGroupByCommonOrder(o *unitOrder) arrayOfUnits {
	var arr []*unit
	for _, u := range s.units {
		if u.currentOrder == o {
			arr = append(arr, u)
		}
	}
	return arr
}

func (s *scene) performOrdersForUnits() {
	for _, u := range s.units {
		if u.currentOrder == nil || u.movementPointsRemaining == 0 {
			continue
		}
		ug := s.getUnitGroupByCommonOrder(u.currentOrder)
		if ug.getMinMovementPoints() == 0 {
			continue
		}
		switch u.currentOrder.orderCode {
		case ORDER_MOVE:
			s.performMoveOrderForUnits(ug)
		case ORDER_HARVEST:
			s.performHarvestOrderForUnits(ug)
		}
	}
}

func (s *scene) performMoveOrderForUnits(unts arrayOfUnits) {
	ux, uy := unts.getCoords()
	tx, ty := unts[0].currentOrder.getCoords()
	vx, vy := s.getVectorForPathFromTo(ux, uy, tx, ty)
	if !s.tryImmediateMoveUnits(unts, vx, vy) {
		unts.emptyActionPoints()
	}
	ux, uy = unts.getCoords()
	if ux == tx && uy == ty {
		unts.resetOrder()
	}
}

func (s *scene) performHarvestOrderForUnits(unts arrayOfUnits) {
	ux, uy := unts.getCoords()
	tx, ty := unts[0].currentOrder.getCoords()

	// find new coords for harvest
	if s.tiles[tx][ty].resourceAmountHere == 0 {
		tx, ty = calculations.SpiralSearchForClosestConditionFrom(func(x, y int) bool {
			return s.areCoordsValid(x, y) && s.tiles[x][y].resourceAmountHere > 0
		}, ux, uy, len(s.tiles), 0)
		if tx == -1 {
			unts[0].owner.addNotification("Nowhere to harvest.")
			unts.resetOrder()
			return
		}
		unts[0].currentOrder.x, unts[0].currentOrder.y = tx, ty
	}
	// 1. Set move coords at resources
	moveToX, moveToY := tx, ty
	// 2. Some Harvesters are full, we are moving to the city
	if unts.areSomeHarvestersWithCargo() {
		city := s.getCityAcceptingHarvestersClosestTo(unts[0].owner, ux, uy)
		if city == nil {
			unts.emptyActionPoints()
			return
		} else {
			moveToX, moveToY = city.x, city.y
		}
		if ux == moveToX && uy == moveToY {
			// 3. We're at the city, unload
			for _, u := range unts {
				if u.canHarvest() {
					u.owner.currentMoney += u.currentHarvested
					u.currentHarvested = 0
				}
			}
			return
		}
	}
	// 4. Harvesters are not full, we're already at the resources
	if !unts.areAllHarvestersFull() && ux == moveToX && uy == moveToY {
		if unts.getMinMovementPoints() > 0 {
			for _, u := range unts {
				if u.canHarvest() {
					harvested := calculations.MinInt(u.remainingResourceCapacity(), s.tiles[ux][uy].resourceAmountHere)
					s.tiles[ux][uy].resourceAmountHere -= harvested
					u.currentHarvested += harvested
				}
			}
			unts.emptyActionPoints()
		}
		return
	}

	// move
	vx, vy := s.getVectorForPathFromTo(ux, uy, moveToX, moveToY)
	if !s.tryImmediateMoveUnits(unts, vx, vy) {
		unts.emptyActionPoints()
	}
}

func (s *scene) tryImmediateMoveUnits(unts arrayOfUnits, vx, vy int) bool {
	cx, cy := unts.getCoords()
	if !s.areCoordsValid(cx+vx, cy+vy) {
		return false
	}
	if unts.getMinMovementPoints() > 0 {
		unts.moveAllByVector(vx, vy)
		s.performExploration()
		unts.spendMovementPoints(1)
		return true
	}
	return false
}
