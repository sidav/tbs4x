package main

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
			return
		}
		ug := s.getUnitGroupByCommonOrder(u.currentOrder)
		if ug.getMinMovementPoints() == 0 {
			return
		}
		switch u.currentOrder.orderCode {
		case ORDER_MOVE:
			s.performMoveOrderForUnits(ug)
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
