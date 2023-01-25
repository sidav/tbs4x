package main

type unitStaticData struct {
	name            string
	figuresInUnit   int
	healthPerFigure int

	geoscapeStats unitStaticGeoscapeStats
}

type unitStaticGeoscapeStats struct {
	speed  int
	vision int
}

const (
	UNT_RECON = iota
)

var sTableUnits = map[int]*unitStaticData{
	UNT_RECON: {
		name:            "Recon squad",
		figuresInUnit:   2,
		healthPerFigure: 2,
		geoscapeStats: unitStaticGeoscapeStats{
			speed:  2,
			vision: 2,
		},
	},
}
