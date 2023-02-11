package main

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

func (units arrayOfUnits) moveAllByVector(vx, vy int) {
	for _, u := range units {
		u.x += vx
		u.y += vy
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
