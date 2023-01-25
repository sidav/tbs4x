package main

type unit struct {
	owner *player
	code  int
	x, y  int
}

func (u *unit) getStaticData() *unitStaticData {
	return sTableUnits[u.code]
}
