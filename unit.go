package main

type unit struct {
	owner                   *player
	id                      int
	x, y                    int
	movementPointsRemaining int

	currentHarvested int

	currentOrder *unitOrder
}

func (u *unit) initByStatic() {
	u.movementPointsRemaining = u.getStaticData().geoscapeStats.speed
}

func (u *unit) getCoords() (int, int) {
	return u.x, u.y
}

func (u *unit) getStaticData() *unitStaticData {
	return sTableUnits[u.id]
}

func (u *unit) canHarvest() bool {
	return u.getStaticData().geoscapeStats.resourceCapacity > 0
}

func (u *unit) remainingResourceCapacity() int {
	return u.getStaticData().geoscapeStats.resourceCapacity - u.currentHarvested
}

func (u *unit) canPerformOrder(code int) bool {
	switch code {
	case ORDER_NONE, ORDER_MOVE, ORDER_EXPLORE:
		return true
	case ORDER_HARVEST:
		return u.canHarvest()
	default:
		return false
	}
}
