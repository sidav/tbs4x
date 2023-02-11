package main

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
