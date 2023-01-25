package main

type unit struct {
	code int
	x, y int
}

func (u *unit) getStaticData() *unitStaticData {
	return sTableUnits[u.code]
}
