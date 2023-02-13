package main

type unit struct {
	owner                   *player
	id                      int
	x, y                    int
	movementPointsRemaining int

	currentOrder unitOrder
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

func (u *unit) canPerformOrder(code int) bool {
	switch code {
	case ORDER_NONE, ORDER_MOVE:
		return true
	default:
		return false
	}
}
