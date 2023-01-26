package main

func (s *scene) tryImmediateMoveUnits(unts []*unit, vx, vy int) bool {
	for _, u := range unts {
		u.x += vx
		u.y += vy
	}
	s.performExploration()
	return true
}
