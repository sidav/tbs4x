package main

type unit struct {
	owner                   *player
	code                    int
	x, y                    int
	movementPointsRemaining int
}

func (u *unit) initByStatic() {
	u.movementPointsRemaining = u.getStaticData().geoscapeStats.speed
}

func (u *unit) getCoords() (int, int) {
	return u.x, u.y
}

func (u *unit) getStaticData() *unitStaticData {
	return sTableUnits[u.code]
}
