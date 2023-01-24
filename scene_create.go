package main

import "tbs4x/lib/perlin_noise"

func (s *scene) init() {
	// Perlin noise generation
	noiseMap := perlin_noise.GeneratePerlinNoiseMap(64, 64, 20, 8, rnd)
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
			case noiseMap[x][y] < 0.55:
				s.tiles[x][y] = &tile{code: TILE_SAND}
			case noiseMap[x][y] < 0.90:
				s.tiles[x][y] = &tile{code: TILE_GRASS}
			default:
				s.tiles[x][y] = &tile{code: TILE_MOUNTAIN}
			}
		}
	}
	s.generateResources()

	x, y := -1, -1
	for !s.isTileApplicableForCity(x, y) {
		x, y = rnd.RandInRange(0, 64), rnd.RandInRange(0, 64)
	}
	s.addCity(&city{
		name: "Alpha Base",
		x:    x,
		y:    y,
	})
}

func (s *scene) generateResources() {
	var desiredTotalRes = map[int]int{
		RES_NONE:     0,
		RES_GOLD:     1000,
		RES_GREENIUM: 1000,
	}
	currTotalRes := make(map[int]int, 0)
	for res := 1; res <= 2; res++ {
		for currTotalRes[res] < desiredTotalRes[res] {
			x, y := -1, -1
			for !s.areCoordsValid(x, y) || s.tiles[x][y].resourceCode != RES_NONE ||
				!s.tiles[x][y].getStaticData().canHaveResource(res) {

				x, y = rnd.RandInRange(0, 64), rnd.RandInRange(0, 64)
			}
			s.tiles[x][y].resourceCode = res
			resAmount := rnd.RandInRange(50, 150)
			s.tiles[x][y].resourceAmountHere = resAmount
			currTotalRes[res] += resAmount
		}
	}
}
