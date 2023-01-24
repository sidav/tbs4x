package main

type tile struct {
	code               int
	resourceCode       int
	resourceAmountHere int
}

func (t *tile) getStaticData() *tileStaticData {
	return sTableTiles[t.code]
}

const (
	RES_NONE = iota
	RES_GOLD
	RES_GREENIUM
)

const (
	TILE_SAND = iota
	TILE_GRASS
	TILE_MOUNTAIN
	TILE_WATER
)

type tileStaticData struct {
	possibleResources []int
}

func (tsd *tileStaticData) canHaveResource(code int) bool {
	for _, r := range tsd.possibleResources {
		if r == code {
			return true
		}
	}
	return false
}

var sTableTiles = map[int]*tileStaticData{
	TILE_SAND: {
		possibleResources: []int{RES_GREENIUM},
	},
	TILE_MOUNTAIN: {},
	TILE_WATER:    {},
	TILE_GRASS:    {possibleResources: []int{RES_GOLD}},
}
