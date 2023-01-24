package main

type tile struct {
	code int
}

func (t *tile) getStaticData() *tileStaticData {
	return sTableTiles[t.code]
}

const (
	TILE_SAND = iota
	TILE_ROCK
	TILE_WATER
)

type tileStaticData struct {
}

var sTableTiles = map[int]*tileStaticData{
	TILE_SAND:  {},
	TILE_ROCK:  {},
	TILE_WATER: {},
}
