package main

type unit struct {
	owner *player
	code  int
	x, y  int
}

func (u *unit) getCoords() (int, int) {
	return u.x, u.y
}

func (u *unit) getStaticData() *unitStaticData {
	return sTableUnits[u.code]
}
