package main

import "tbs4x/lib/perlin_noise"

func (s *scene) init(w, h int) {
	s.currentTurn = 1
	for i := 0; i < 100; i++ {
		// Perlin noise generation
		noiseMap := perlin_noise.GeneratePerlinNoiseMap(w, h, 16, 8, rnd)
		perlin_noise.MakeNoiseRectangular(noiseMap)
		s.tiles = make([][]*tile, len(noiseMap))
		for i := range s.tiles {
			s.tiles[i] = make([]*tile, len(noiseMap[i]))
		}
		for x := range noiseMap {
			for y := range noiseMap[x] {
				switch {
				case noiseMap[x][y] < 0.5:
					s.tiles[x][y] = &tile{code: TILE_WATER}
				case noiseMap[x][y] < 0.575:
					s.tiles[x][y] = &tile{code: TILE_SAND}
				case noiseMap[x][y] < 0.90:
					s.tiles[x][y] = &tile{code: TILE_GRASS}
				default:
					s.tiles[x][y] = &tile{code: TILE_MOUNTAIN}
				}
			}
		}
		s.generateResources()

		if s.tryAddPlayer() {
			return
		}
	}
	panic("Player is unplaceable")
}

func (s *scene) tryAddPlayer() bool {
	newPlayer := &player{}
	newPlayer.init(s.getSize())
	x, y := -1, -1
	try := 0
	for !s.isTileApplicableForCity(x, y) {
		x, y = s.getRandomCoords()
		try++
		if try > 10000 {
			return false
		}
	}
	startCity := &city{
		owner:        newPlayer,
		name:         "Alpha Base",
		x:            x,
		y:            y,
		maxBuildings: s.countTilesAllowingBuildingAround(x, y, 2),
	}
	startCity.addBuilding(findBuildingInTableByName("hq"))
	s.addCity(startCity)
	s.addUnit(findUnitStaticIndexByName("recon"), newPlayer, x+1, y+1)
	s.addUnit(findUnitStaticIndexByName("harvester"), newPlayer, x, y+1)
	s.players = append(s.players, newPlayer)
	return true
}

func (s *scene) generateResources() {
	const minInPatch, maxInPatch = 50, 350
	var desiredTotalRes = map[int]int{
		RES_NONE: 0,
		RES_GOLD: 5000,
		// RES_GREENIUM: 1000,
	}
	currTotalRes := make(map[int]int, 0)
	for res := 1; res <= 2; res++ {
		for currTotalRes[res] < desiredTotalRes[res] {
			x, y := -1, -1
			try := 0
			for !s.areCoordsValid(x, y) || s.tiles[x][y].resourceCode != RES_NONE ||
				!s.tiles[x][y].getStaticData().canHaveResource(res) {

				x, y = s.getRandomCoords()
				try++
				if try > 10000 {
					panic("Can't place resource")
				}
			}
			s.tiles[x][y].resourceCode = res
			resAmount := rnd.RandInRange(minInPatch, maxInPatch)
			s.tiles[x][y].resourceAmountHere = resAmount
			currTotalRes[res] += resAmount
		}
	}
}
