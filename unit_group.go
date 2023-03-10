package main

import "fmt"

type arrayOfUnits []*unit

func (units arrayOfUnits) getCoords() (int, int) {
	x, y := units[0].getCoords()
	for _, u := range units {
		if u.x != x || u.y != y {
			panic("Selected units are not in squad!")
		}
	}
	return x, y
}

func (units arrayOfUnits) getOrder() *unitOrder {
	order := units[0].currentOrder
	for _, u := range units {
		if u.currentOrder != order {
			panic("Many orders in single group!")
		}
	}
	return order
}

func (units arrayOfUnits) getOwner() *player {
	owner := units[0].owner
	for _, u := range units {
		if u.owner != owner {
			panic("Many owners in single group!")
		}
	}
	return owner
}

func (units arrayOfUnits) notifyOwner(note string) {
	x, y := units.getCoords()
	units.getOwner().addNotification(fmt.Sprintf("Squad at %d,%d: %s", x, y, note))
}

func (units arrayOfUnits) moveAllByVector(vx, vy int) {
	for _, u := range units {
		u.x += vx
		u.y += vy
	}
}

func (units arrayOfUnits) areSomeHarvestersWithCargo() bool {
	for _, u := range units {
		if u.canHarvest() && u.currentHarvested > 0 {
			return true
		}
	}
	return false
}

func (units arrayOfUnits) areAllHarvestersFull() bool {
	for _, u := range units {
		if u.canHarvest() && u.currentHarvested < u.getStaticData().geoscapeStats.resourceCapacity {
			return false
		}
	}
	return true
}

func (units arrayOfUnits) resetOrder() {
	for _, u := range units {
		u.currentOrder = nil
	}
}

func (units arrayOfUnits) emptyActionPoints() {
	for _, u := range units {
		u.movementPointsRemaining = 0
	}
}

func (units arrayOfUnits) spendMovementPoints(pts int) {
	for _, u := range units {
		u.movementPointsRemaining -= pts
	}
}

func (units arrayOfUnits) getMinMovementPoints() int {
	min := -1
	for _, u := range units {
		if min == -1 || u.movementPointsRemaining < min {
			min = u.movementPointsRemaining
		}
	}
	return min
}

func (units arrayOfUnits) canGroupPerformOrder(code int) bool {
	for _, u := range units {
		if u.canPerformOrder(code) {
			return true
		}
	}
	return false
}

func (units arrayOfUnits) assignOrder(o *unitOrder) {
	for _, u := range units {
		u.currentOrder = o
	}
}

//func (units arrayOfUnits) getAvailableOrderCodes() []int {
//	var codes []int
//	for _, u := range units {
//		for currCode := 0;  currCode < ORDERS_COUNT; currCode++ {
//			isInArray := false
//			for i := range codes {
//				if codes[i] == currCode {
//					isInArray = true
//					break
//				}
//			}
//			if !isInArray && u.canPerformOrder(currCode) {
//				codes = append(codes, currCode)
//			}
//		}
//	}
//	return codes
//}
